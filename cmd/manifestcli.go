package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/aleosiss/manifest/cmd/manifest"
	"github.com/aleosiss/manifest/internal/resource"
	"github.com/aleosiss/manifest/internal/util"
	"github.com/aleosiss/manifest/internal/web"
	"os"
	"strings"
	"sync"
)

func main() {
	initialize()

	args, err := handleArguments()
	util.HandleError(err)

	manifestFile, err := manifest.From(args[0])
	util.HandleError(err)
	err = manifest.Validate(manifestFile)

	var files []string

	fmt.Println("Handling Manifest: " + manifestFile.Name)
	wg := sync.WaitGroup{}
	for _, target := range manifestFile.Targets {
		wg.Add(1)
		go func(target manifest.Target) {
			defer wg.Done()
			file := handleTarget(target)
			files = append(files, file)
		} (target)
	}
	wg.Wait()

	err = packageForDeployment(manifestFile.Package.Type, manifestFile.Package.Location, files)
	util.HandleError(err)

	cleanup()
}

func handleArguments() (args []string, err error) {
	flag.Parse()
	args = flag.Args()


	if len(args) < 1 {
		err = errors.New("no manifest file was provided")
	}
	return
}

func handleTarget(target manifest.Target) string {
	fmt.Println("Found target: " + target.Name)
	url, err := util.ExpandText(target.URL, "version", target.TargetVersion)
	util.HandleError(err)

	downloadedTarget, err := web.DownloadTarget(target, url)
	util.HandleError(err)

	processedTarget, err := postprocessTarget(target.PostProcess, downloadedTarget)
	util.HandleError(err)

	return processedTarget
}

func initialize() {
	resource.CreateDirectories()
}

func postprocessTarget(process string, target string) (filePath string, err error) {
	filePath = target
	return
}

func packageForDeployment(packageType manifest.PackageType, location string, files []string) (err error) {
	var archive string

	if packageType == manifest.ZIP {
		archive, err = util.ArchiveZip(files)
		if err != nil && ! strings.Contains(err.Error(), "file already exists") {
			util.HandleError(err)
		}
	} else {
		util.HandleError(errors.New("package instructions did not contain a supported type"))
	}

	err = os.MkdirAll(location, os.ModeDir)
	util.HandleError(err)

	if util.Exists(archive) {
		err := util.MoveFile(archive, location)
		util.HandleError(err)
	} else {
		fmt.Println("archive did not exist")
	}


	return nil
}

func cleanup() {
	util.CleanUp()
}

