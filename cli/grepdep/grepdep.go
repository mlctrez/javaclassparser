package main

import (
	"fmt"
	"strings"

	"github.com/mlctrez/javaclassparser/parser"
)

func main() {
	config := parser.NewConfigFromArgs()

	// com.foo.Bar -> com/foo/Bar
	if strings.Contains(config.MatchTarget, ".") {
		config.MatchTarget = strings.Replace(config.MatchTarget, ".", "/", -1)
	}

	parser.Scan(config, func(work *parser.Work) {

		matched := make(map[string]bool)

		work.Class.RefVisit(func(className string) {

			n := parser.ExtractName(work.Class.Name)

			if config.MatchTarget != "" && !strings.HasPrefix(className, config.MatchTarget) {
				return
			}

			key := fmt.Sprintf("%-100s %s",n,className)

			if !matched[key] {
				matched[key] = true
				fmt.Println(key)
			}

		})
	})
}
