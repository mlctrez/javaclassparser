package parser

import (
	"flag"
	"fmt"
	"os"
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
	MatchTarget     string
}

func NewConfigFromArgs() (config *Config) {
	config = &Config{}

	flag.StringVar(&config.Archive, "archive", "", "the war, jar or ear archive to scan.")
	flag.StringVar(&config.Class, "class", "", "only display information about this class")
	flag.BoolVar(&config.PrintArchives, "pa", false, "print each archive name as it is read")
	flag.BoolVar(&config.PrintClassNames, "pc", false, "print each class name as it is read")
	flag.BoolVar(&config.PrintMethodRef, "pmr", false, "print method ref information")
	flag.BoolVar(&config.LogElapsed, "le", true, "log total elapsed time")
	flag.StringVar(&config.DebugClass, "dbc", "", "dump detailed byte code information for this class")
	flag.StringVar(&config.MatchTarget, "match", "", "field or method reference must match this prefix")

	flag.Parse()

	if config.Archive == "" {
		fmt.Println("archive is required")
		os.Exit(1)
	}
	if f, err := os.Stat(config.Archive); err == nil {
		if f.IsDir() {

		}
	}

	return config
}
