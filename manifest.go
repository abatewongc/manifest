package main

import (
	"flag"
	"github.com/aleosiss/manifest/cmd/manifest"
	"github.com/aleosiss/manifest/internal/globals"
	"github.com/aleosiss/manifest/internal/resource"
	"github.com/aleosiss/manifest/internal/ui"
	"github.com/aleosiss/manifest/internal/util"
	"log"
	"os"
)

var UIMode = false

func main() {
	initialize()

	var err error
	// store cwd for cwd drift issue
	globals.CWD, err = os.Getwd()

	args, err := handleArguments()
	util.HandleError(err)
	var fileErrors []error

	if !UIMode {
		err, fileErrors = manifest.Process(args[0], false)
		for _, err := range fileErrors {
			log.Println(err)
		}
	} else {
		ui.Start()
	}

	util.CleanUp()
}

func initialize() {
	resource.CreateDirectories()
}

func handleArguments() (args []string, err error) {
	flag.Parse()
	args = flag.Args()

	if len(args) < 1 {
		UIMode = true
	}
	return
}

