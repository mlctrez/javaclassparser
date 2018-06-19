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
	"strings"
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
			log.Println(time.Since(start).Seconds())
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
	read(config, config.archive, rc)

}

func read(config *ParserConfig, path string, rc *zip.ReadCloser) {
	if config.printArchives {
		fmt.Println("reading", path)
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
			read(config, path+"!"+f.Name, jarReader)
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
			jcp := &javaclassparser.ClassParser{}
			jcp.Parse(bytes.NewReader(bb.Bytes()))
			if "all" == config.debugClass || config.debugClass == f.Name {
				jcp.DebugOut()
			}

		}
		rp.Close()

	}

}
