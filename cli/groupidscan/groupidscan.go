package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mlctrez/javaclassparser/ioutil"
)

func main() {

	var directory string
	var groupIds string
	flag.StringVar(&directory, "directory", "", "directory containing the jar files")
	flag.StringVar(&groupIds, "groupIds", "", "groupIds to match, comma separated")
	flag.Parse()

	if directory == "" {
		log.Fatal("directory is a required parameter")
	}
	if groupIds == "" {
		log.Fatal("groupIds is a required parameter")
	}

	gids := strings.Split(groupIds, ",")

	filepath.Walk("/Users/mattman/work/ear/APP-INF/lib", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || !strings.HasSuffix(path, ".jar") {
			return nil
		}
		rc, err := zip.OpenReader(path)
		if err != nil {
			panic(err)
		}
		defer rc.Close()
		for _, f := range rc.File {
			if f.FileInfo().IsDir() {
				continue
			}
			if strings.HasPrefix(f.Name, "META-INF") && strings.HasSuffix(f.Name, "pom.properties") {
				var groupId string
				ioutil.ScanLines(ioutil.MustOpen(f.Open()), func(line string) {
					if strings.HasPrefix(line, "groupId=") {
						if groupId != "" {
							panic("found more than one groupId")
						}
						groupId = line[len("groupId="):]
					}
				})
				for _, g := range gids {
					if strings.HasPrefix(groupId, g) {
						fmt.Println(path)
					}
				}
			}
		}
		return nil
	})

}
