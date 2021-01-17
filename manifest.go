package main

import (
	"flag"
	"github.com/aleosiss/manifest/internal/globals"
	"github.com/aleosiss/manifest/internal/resource"
	"github.com/aleosiss/manifest/internal/service"
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

	manifestService := service.NewManifestService()

	if !UIMode {
		err, fileErrors = manifestService.Process(args[0], false)
		for _, err := range fileErrors {
			log.Println(err)
		}
	} else {
		gui := ui.New(&manifestService)
		gui.Start()
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

	if len(args) > 1  {
		log.Println("WARNING: Arguments after first ignored! First argument was " + args[0])
	}
	return
}

