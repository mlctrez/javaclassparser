package parser

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/mlctrez/javaclassparser/attribute"
)

// Config contains various flags used while parsing
type Config struct {
	Archive         string
	Class           string
	PrintArchives   bool
	PrintClassNames bool
	PrintMethodRef  bool
	LogElapsed      bool
	DebugClass      string
	Workers         int
}

// Context manages the input and output work channels and wait groups
type Context struct {
	config      *Config
	work        chan *Work
	workGroup   *sync.WaitGroup
	result      chan *Work
	resultGroup *sync.WaitGroup
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
	for w := range pc.work {
		w.Class, w.Error = New(w.ByteCode)
		pc.resultGroup.Add(1)
		pc.result <- w
		pc.workGroup.Done()
	}
}

// Parse reads all class files in the provided archive and executes the callback for each
// Order is not guaranteed.
func Parse(config *Config, rc *zip.ReadCloser, callback func(*Work)) {

	if config.Workers == 0 {
		config.Workers = runtime.NumCPU()
	}

	pc := &Context{
		config:      config,
		work:        make(chan *Work),
		workGroup:   &sync.WaitGroup{},
		result:      make(chan *Work),
		resultGroup: &sync.WaitGroup{},
	}

	for i := 0; i < config.Workers; i++ {
		go parseWorker(pc)
	}

	go func() {
		for w := range pc.result {
			callback(w)
			pc.resultGroup.Done()
		}
	}()
	readArchive(pc, config.Archive, rc)

	pc.workGroup.Wait()
	close(pc.work)

	pc.resultGroup.Wait()
	close(pc.result)
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

			pc.workGroup.Add(1)
			pc.work <- &Work{Path: path, ByteCode: bb, Config: pc.config}

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
