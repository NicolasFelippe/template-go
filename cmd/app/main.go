package main

import (
	"flag"
	"fmt"
	_ "github.com/golang/mock/mockgen/model"
	"os"
	project "template-go/cmd/app/internal"
	log "template-go/internal/logger"
)

var (
	buildVersion string
	buildCommit  string
)

const (
	usage = `Usage:
	backend -config_file=./configs/app.env
	`
)

func main() {
	exitCode := 0
	f := project.Flags{}

	flag.Usage = func() {
		_, err := fmt.Fprintf(os.Stderr, "%s\n", usage)
		if err != nil {
			return
		}
	}
	flag.StringVar(&f.ConfigurationFile, "config_file", "./configs",
		"configuration file of the project")
	flag.StringVar(&f.TimeZone, "tz", "", "time zone of the project")
	flag.BoolVar(&f.Version, "version", false, "to print version of the program")
	flag.Var(&f.Debug, "debug", "turn on debug mode, this will set log level to debug")
	flag.Parse()

	if f.Version {
		_, err := fmt.Fprintf(os.Stderr, "version: %s\ncommit: %s\n", buildVersion, buildCommit)
		if err != nil {
			return
		}

		return
	}

	if err := project.Run(f); err != nil {
		log.Logger.Fatal(fmt.Sprintf("Fatal Error: %v", err))
		exitCode = 1
	}

	os.Exit(exitCode)
}
