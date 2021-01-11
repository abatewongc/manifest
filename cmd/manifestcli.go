package main

import (
	"flag"
	"github.com/aleosiss/manifest/cmd/manifest"
	"github.com/aleosiss/manifest/internal/resource"
	"github.com/aleosiss/manifest/internal/ui"
	"github.com/aleosiss/manifest/internal/util"
)

var UIMode = false

func main() {
	initialize()
	args, err := handleArguments()
	util.HandleError(err)
	if !UIMode {
		manifest.Manifesto(args[0], false)
	} else {
		ui.Start()
	}

	util.CleanUp()
}

func handleArguments() (args []string, err error) {
	flag.Parse()
	args = flag.Args()


	if len(args) < 1 {
		UIMode = true
	}
	return
}


func initialize() {
	resource.CreateDirectories()
}

