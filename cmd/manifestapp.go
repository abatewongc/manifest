package main

import (
	"fmt"
	"github.com/aleosiss/manifest/cmd/manifest"
	"github.com/aleosiss/manifest/internal/resource"
	"github.com/aleosiss/manifest/internal/util"
	"github.com/aleosiss/manifest/internal/web"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	initialize()

	manifestFile, err := manifest.From("./test/examples/manifest.json")
	if err != nil {
		fmt.Println("Error " + err.Error())
		return
	}

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
	HandleError(err)

	cleanup()
}

func handleTarget(target manifest.Target) string {
	fmt.Println("Found target: " + target.Name)
	url, err := formatURLWithVersion(target)
	HandleError(err)

	downloadedTarget, err := downloadTarget(target, url)
	HandleError(err)

	processedTarget, err := postprocessTarget(target.PostProcess, downloadedTarget)
	HandleError(err)

	return processedTarget
}

func cleanup() {
	//util.CleanUp()
}

func HandleError(err error) {
	if err != nil { panic(err) }
}

func initialize() {
	resource.CreateDirectories()
}

func packageForDeployment(packageType manifest.PackageType, location string, files []string) (err error) {
	var archive string

	if packageType == manifest.ZIP {
		archive, err = util.ArchiveZip(files)
		if err != nil && ! strings.Contains(err.Error(), "file already exists") {
			HandleError(err)
		}
	}


	err = os.MkdirAll(location, os.ModeDir)
	if err != nil {
		log.Fatal(err)
	}

	if util.Exists(archive) {
		err := util.MoveFile(archive, location)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("archive did not exist")
	}


	return nil
}

func stageTarget(targetPath string, files []string) []string {
	files = append(files, targetPath)
	return files
}

func postprocessTarget(process string, target string) (filePath string, err error) {
	filePath = target

	return
}

func downloadTarget(target manifest.Target, url string) (filePath string, err error) {
	return web.DownloadTarget(target, url)
}

func formatURLWithVersion(target manifest.Target) (string, error) {
	return util.ExpandText(target.URL, "version", target.TargetVersion)
}
