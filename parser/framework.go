package parser

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/mlctrez/javaclassparser/attribute"
)

// Context manages the input and output work channels and wait groups
type Context struct {
	config *Config
	work   *WorkChanGroup
	result *WorkChanGroup
}

// Work is the communication to and from the workers
type Work struct {
	Config   *Config
	Path     string
	ByteCode *bytes.Buffer
	Class    *Class
	Error    error
}

// DefaultSort sorts a slice of Work by Path then Class.Name
/*

	Example for sorting Work output

	var results []*Work
	sort.SliceStable(results, DefaultSort(results))
*/
func DefaultSort(results []*Work) func(i, j int) bool {
	return func(i, j int) bool {
		if results[i].Path != results[j].Path {
			return results[i].Path < results[j].Path
		}
		return results[i].Class.Name < results[j].Class.Name
	}
}

func parseWorker(pc *Context) {
	for w := range pc.work.ch {
		w.Class, w.Error = New(w.ByteCode)
		pc.result.wg.Add(1)
		pc.result.ch <- w
		pc.work.wg.Done()
	}
}

// Scan reads all class files in the provided config.Archive path
// if config.Archive is a directory, all jar files within the directory are scanned
func Scan(config *Config, callback func(*Work)) (err error) {
	if config.Archive == "" {
		return errors.New("archive must be provided")
	}

	var info os.FileInfo
	if info, err = os.Stat(config.Archive); err != nil {
		return err
	}

	if !info.IsDir() {
		var rc *zip.ReadCloser
		if rc, err = zip.OpenReader(config.Archive); err != nil {
			return err
		}
		Parse(config, rc, callback)
		return nil
	}

	originalArchive := config.Archive
	defer func() {
		config.Archive = originalArchive
	}()
	return filepath.Walk(config.Archive, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), ".jar") || strings.HasSuffix(info.Name(), ".ear") || strings.HasSuffix(info.Name(), ".war") {
			var rc *zip.ReadCloser
			if rc, err = zip.OpenReader(path); err != nil {
				return err
			}
			config.Archive = path
			Parse(config, rc, callback)
		}
		return nil
	})
	return
}

type WorkChanGroup struct {
	ch chan *Work
	wg *sync.WaitGroup
}

func newWorkChanGroup(channelSize int) (wc *WorkChanGroup) {
	wc = &WorkChanGroup{}
	wc.ch = make(chan *Work, channelSize)
	wc.wg = &sync.WaitGroup{}
	return wc
}

// Parse reads all class files in the provided archive and executes the callback for each
// Order is not guaranteed.
func Parse(config *Config, rc *zip.ReadCloser, callback func(*Work)) {

	start := time.Now()

	if config.Workers == 0 {
		config.Workers = runtime.NumCPU()
	}

	channelSize := config.Workers * 2

	pc := &Context{
		config: config,
		work:   newWorkChanGroup(channelSize),
		result: newWorkChanGroup(channelSize),
		//work:        make(chan *Work, channelSize),
		//workGroup:   &sync.WaitGroup{},
		//result:      make(chan *Work, channelSize),
		//resultGroup: &sync.WaitGroup{},
	}

	for i := 0; i < config.Workers; i++ {
		go parseWorker(pc)
	}

	var total uint16

	go func() {
		for w := range pc.result.ch {
			total++
			callback(w)
			pc.result.wg.Done()
		}
	}()
	readArchive(pc, config.Archive, rc)

	pc.work.wg.Wait()
	close(pc.work.ch)

	pc.result.wg.Wait()
	close(pc.result.ch)

	if config.LogElapsed {
		fmt.Printf("%05d classes scanned in %08.5f sec from %s\n",
			total, time.Since(start).Seconds(), config.Archive)
	}
}

func readArchive(pc *Context, path string, rc *zip.ReadCloser) {
	if pc.config.PrintArchives {
		fmt.Println("reading", path)
	}

	for _, f := range rc.File {
		if f.FileInfo().IsDir() {
			continue
		}
		rp, err := f.Open()
		if err != nil {
			// TODO: handle this better
			panic(err)
		}

		if strings.HasSuffix(f.Name, ".jar") {

			tf, err := ioutil.TempFile(os.TempDir(), "jcp")
			if err != nil {
				panic(err)
			}
			_, err = io.Copy(tf, rp)
			if err != nil {
				panic(err)
			}
			jarReader, err := zip.OpenReader(tf.Name())
			if err != nil {
				panic(err)
			}

			readArchive(pc, path+"!"+f.Name, jarReader)
			jarReader.Close()
			os.Remove(tf.Name())
		}

		if strings.HasSuffix(f.Name, ".class") {

			bb := &bytes.Buffer{}
			bc, err := io.Copy(bb, rp)
			if err != nil {
				log.Fatal(err)
			}

			if bc != int64(f.UncompressedSize64) {
				fmt.Println(bc, f.UncompressedSize64, f.Name, err)
				log.Fatal("unable to read entire file")
			}
			if pc.config.PrintClassNames {
				fmt.Println("reading", f.Name)
			}

			pc.work.wg.Add(1)
			pc.work.ch <- &Work{Path: path, ByteCode: bb, Config: pc.config}

		}
		rp.Close()
	}
}

func (jcp *Class) DebugOut() {

	jcp.pool.DebugOut()

	fmt.Print("access, className, superClass = ")
	fmt.Println(jcp.accessFlags, jcp.pool.Lookup(jcp.classNameIndex), jcp.pool.Lookup(jcp.superClassNameIndex))

	for i, itf := range jcp.interfaces {
		fmt.Printf("interface %3d %s\n", i, jcp.pool.Lookup(itf))
	}

	fmt.Println("*** class fields")

	for i, f := range jcp.fields {
		fmt.Println(i, f)
	}

	fmt.Println("*** class methods")

	for i, f := range jcp.methods {
		fmt.Println(i, f)
		for i := 0; i < len(f.Attributes); i++ {
			attr := f.Attributes[i]
			if code, ok := attr.(*attribute.CodeAttribute); ok {
				for j := 0; j < len(code.Code); j++ {
					instruction := code.Code[j]
					fmt.Printf(" Code %04X %s\n", instruction.Offset, instruction.StringWithIndex(jcp.pool))
				}
				for j := 0; j < len(code.ExceptionTable); j++ {
					fmt.Printf(" ExceptionTable %+v\n", code.ExceptionTable[j])
				}
				for j := 0; j < len(code.Attributes); j++ {
					fmt.Printf(" Attributes %+v\n", code.Attributes[j])
				}
			} else {
				fmt.Printf(" attribute %d %+v\n", i, f.Attributes[i])
			}
		}
	}

	fmt.Println("*** class attributes")

	for i, f := range jcp.attributes {
		fmt.Println(i, f)
	}

}
