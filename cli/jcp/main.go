package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/mlctrez/javaclassparser/cpool"
	"github.com/mlctrez/javaclassparser/parser"
)

func failErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func NewConfigFromArgs() *parser.Config {
	config := &parser.Config{}

	flag.StringVar(&config.Archive, "archive", "", "the war, jar or ear archive to scan")
	flag.StringVar(&config.Class, "class", "", "only display information about this class")
	flag.BoolVar(&config.PrintArchives, "pa", false, "print each archive name as it is read")
	flag.BoolVar(&config.PrintClassNames, "pc", false, "print each class name as it is read")
	flag.BoolVar(&config.PrintMethodRef, "pmr", false, "print method ref information")
	flag.BoolVar(&config.LogElapsed, "le", true, "log total elapsed time")
	flag.StringVar(&config.DebugClass, "dbc", "", "dump detailed byte code information for this class")

	flag.Parse()

	if config.Archive == "" {
		fmt.Println("archive is required")
		os.Exit(1)
	}
	return config
}

func main() {

	//log.SetOutput(os.Stdout)

	config := NewConfigFromArgs()

	results := make([]*parser.Work, 0)
	parser.Scan(config, func(work *parser.Work) {
		results = append(results, work)
	})

	sort.SliceStable(results, parser.DefaultSort(results))

	prefixes := []string{"com/"}
	for _, r := range results {

		// TODO: should one class error out the whole thing?
		if r.Error != nil {
			cn := "<unknown>"
			if r.Class != nil && r.Class.Name != "" {
				cn = r.Class.Name
			}
			log.Fatal(r.Path, cn, r.Error)
		}

		if "all" == config.DebugClass || config.DebugClass == r.Class.Name {
			r.Class.DebugOut()
		}
		if config.PrintMethodRef {

			cn := r.Class.Name
			wanted := false
			for _, p := range prefixes {
				if strings.HasPrefix(cn, p) {
					wanted = true
				}
			}
			if !wanted {
				continue
			}

			r.Class.Visit(func(i interface{}) {
				var ref string
				if mr, ok := i.(*cpool.ConstantMethodrefInfo); ok {
					ref = mr.String()
				}
				if mr, ok := i.(*cpool.ConstantInterfaceMethodrefInfo); ok {
					ref = mr.String()
				}
				if ref == "" || strings.HasPrefix(ref, "java/") || strings.HasPrefix(ref, "org/") || strings.HasPrefix(ref, "javax/") {
					return
				}
				if strings.HasPrefix(ref, cn) || strings.HasPrefix(ref, "[L"+cn+";") {
					return
				}

				fmt.Println(cn, "->", ref)
			})

		}
	}

}
