package main

import (
	"fmt"
	"github.com/aleosiss/manifest/cmd/manifest"
	"github.com/aleosiss/manifest/internal/resource"
	"github.com/aleosiss/manifest/internal/util"
	"github.com/aleosiss/manifest/internal/web"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	initialize()

	manifest, err := manifest.From("./test/examples/manifest.json")
	if err != nil {
		fmt.Println("Error " + err.Error())
		return
	}

	var files []string

	fmt.Println("Handling Manifest: " + manifest.Name)
	for _, target := range manifest.Targets {
		files = manifesto(target, files)
	}

	err = packageForDeployment(manifest.Package.Type, manifest.Package.Location, files)
	HandleError(err)

	cleanup()
}

func manifesto(target manifest.Target, files []string) []string {
	fmt.Println("Found target: " + target.Name)
	url, err := formatURLWithVersion(target)
	HandleError(err)

	downloadedTarget, err := downloadTarget(target, url)
	HandleError(err)

	processedTarget, err := postprocessTarget(target.PostProcess, downloadedTarget)
	HandleError(err)

	files = stageTarget(processedTarget, files)
	return files
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
		if ! strings.Contains(err.Error(), "file already exists") {
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
	stagedPath := resource.ManifestStagingDir + string(filepath.Separator) + filepath.Base(targetPath)

	err := util.MoveFile(targetPath, stagedPath)
	if err != nil {
		log.Fatal(err)
	}

	files = append(files, stagedPath)
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
