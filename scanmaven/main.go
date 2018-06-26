package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	current, err := user.Current()
	if err != nil {
		panic(err)
	}

	_, err = exec.Command("go", "build", "-o", "jcp_bin", "cli/jcp/main.go").CombinedOutput()
	if err != nil {
		panic(err)
	}

	mavenRepository := path.Join(current.HomeDir, ".m2/repository")
	filepath.Walk(mavenRepository, func(path string, info os.FileInfo, err error) error {

		if strings.Contains(path, "-sources.jar") || strings.Contains(path, "-javadoc.jar") {
			return nil
		}

		if strings.HasSuffix(path, ".jar") {
			bytes, err := exec.Command("./jcp_bin", "-archive", path).CombinedOutput()

			output := strings.TrimSpace(string(bytes))
			if strings.Contains(output, "unhandled attribute") {
				log.Fatal(output)
			}
			fmt.Println(output)
			if err != nil {
				log.Fatal(err)
				return nil
			}
		}
		return nil
	})
}
