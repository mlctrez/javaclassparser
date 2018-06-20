package main

import (
	"archive/zip"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/mlctrez/javaclassparser"
)

func failErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type ParserConfig struct {
	archive         string
	class           string
	printArchives   bool
	printClassNames bool
	logElapsed      bool
	debugClass      string
}

func NewConfigFromArgs() *ParserConfig {
	config := &ParserConfig{}

	flag.StringVar(&config.archive, "archive", "", "the war, jar or ear archive to scan")
	flag.StringVar(&config.class, "class", "", "only display information about this class")
	flag.BoolVar(&config.printArchives, "pa", false, "print each archive name as it is read")
	flag.BoolVar(&config.printClassNames, "pc", false, "print each class name as it is read")
	flag.BoolVar(&config.logElapsed, "le", true, "log total elapsed time")
	flag.StringVar(&config.debugClass, "dbc", "", "dump detailed byte code information for this class")

	flag.Parse()

	if config.archive == "" {
		fmt.Println("archive is required")
		os.Exit(1)
	}
	return config
}

func main() {

	// red is bad
	log.SetOutput(os.Stdout)

	config := NewConfigFromArgs()

	var start time.Time
	if config.logElapsed {
		start = time.Now()
	}
	defer func() {
		if config.logElapsed {
			fmt.Println(time.Since(start).Seconds(), "elapsed seconds")
		}
	}()

	if strings.HasSuffix(config.archive, ".class") {
		jcp := &javaclassparser.ClassParser{}
		b, err := ioutil.ReadFile(config.archive)
		if err != nil {
			panic(err)
		}
		jcp.Parse(bytes.NewReader(b))
		jcp.DebugOut()
		return
	}

	rc, err := zip.OpenReader(config.archive)
	failErr(err)
	parsers := read(config, config.archive, rc)
	fmt.Println("executed", len(parsers), "parsers")
	sort.SliceStable(parsers, func(i, j int) bool {
		if parsers[i].Path != parsers[j].Path {
			return parsers[i].Path < parsers[j].Path
		}
		return parsers[i].Class < parsers[j].Class
	})
	for _, r := range parsers {
		if "all" == config.debugClass || config.debugClass == r.Class {
			r.DebugOut()
		}
	}

}

func read(config *ParserConfig, path string, rc *zip.ReadCloser) []*javaclassparser.ClassParser {
	if config.printArchives {
		fmt.Println("reading", path)
	}

	result := make([]*javaclassparser.ClassParser, 0)

	workWaitGroup := &sync.WaitGroup{}
	workChan := make(chan *work, 50)

	resultChan := make(chan *javaclassparser.ClassParser, 50)
	resultWaitGroup := &sync.WaitGroup{}

	go func() {
		for j := range resultChan {
			result = append(result, j)
			resultWaitGroup.Done()
		}

	}()

	for i := 0; i < runtime.NumCPU(); i++ {
		go parseWorker(workWaitGroup, resultWaitGroup, workChan, resultChan)
	}

	for _, f := range rc.File {
		if f.FileInfo().IsDir() {
			continue
		}
		rp, err := f.Open()
		failErr(err)

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
			result = append(result, read(config, path+"!"+f.Name, jarReader)...)
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
			if config.printClassNames {
				fmt.Println("reading", f.Name)
			}

			workWaitGroup.Add(1)
			workChan <- &work{Class: f.Name, Path: path, ByteCode: bb, config: config}

		}
		rp.Close()
	}
	workWaitGroup.Wait()
	close(workChan)

	resultWaitGroup.Wait()
	close(resultChan)

	return result

}

type work struct {
	Class    string
	Path     string
	ByteCode *bytes.Buffer
	config   *ParserConfig
}

func parseWorker(wg *sync.WaitGroup, rwg *sync.WaitGroup, workChan <-chan *work, resultChan chan<- *javaclassparser.ClassParser) {
	for w := range workChan {
		jcp := &javaclassparser.ClassParser{Class: w.Class, Path: w.Path}
		jcp.Parse(w.ByteCode)
		rwg.Add(1)
		resultChan <- jcp
		wg.Done()
	}

}
