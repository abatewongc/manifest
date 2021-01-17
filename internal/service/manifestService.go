package service

import (
	"errors"
	"fmt"
	gui "github.com/AllenDang/giu"
	"github.com/agrison/go-commons-lang/stringUtils"
	"github.com/aleosiss/manifest/internal/globals"
	"github.com/aleosiss/manifest/internal/model/manifest"
	"github.com/aleosiss/manifest/internal/util"
	"github.com/aleosiss/manifest/internal/web"
	"os"
	"strings"
	"sync"
	"time"
)

type ManifestService struct {

}

func NewManifestService() (manifestService ManifestService) {
	return manifestService
}



func (self *ManifestService) Process(filePath string, uiEnabled bool) (err error, fileErrors []error) {
	globals.UIMsgBoxContents = ""
	if uiEnabled {
		globals.UIProgressBarLabel = "Validating Manifest..."
		gui.Update()
		time.Sleep(1 * time.Second)
	}

	manifestFile, err := manifest.From(filePath)
	if util.HandleError(err) {
		return err, fileErrors
	}

	err = manifestFile.Validate()

	var files []string

	if uiEnabled {
		globals.UIProgressBarLabel = "Downloading Targets..."
		gui.Update()
		time.Sleep(2 * time.Second)
	} else {
		fmt.Println("Handling Manifest: " + manifestFile.Name)
	}

	wg := sync.WaitGroup{}
	for _, target := range manifestFile.Targets {
		wg.Add(1)
		go func(target manifest.Target) {
			defer wg.Done()
			file, err := self.handleTarget(target)

			if err != nil { fileErrors = append(fileErrors, err) }
			if file != "" { files = append(files, file) }
		}(target)
	}
	wg.Wait()

	if uiEnabled {
		globals.UIProgressBarLabel = "Packaging files..."
		gui.Update()
		time.Sleep(1 * time.Second)
	}

	msg, err := self.packageForDeployment(manifestFile.Package.Type, manifestFile.Package.Location, files)
	globals.UIMsgBoxContents = msg
	if util.HandleError(err) {
		return err, fileErrors
	}

	self.cleanup()
	return err, fileErrors
}

func (self *ManifestService) handleTarget(target manifest.Target) (string, error) {
	fmt.Println("Found target: " + target.Name)
	url, err := util.ExpandText(target.URL, "version", target.TargetVersion)
	if util.HandleError(err) {
		return "", err
	}

	downloadedTarget, err := web.DownloadTarget(url)
	if util.HandleError(err) {
		return "", err
	}

	processedTarget, err := self.postprocessTarget(target.PostProcess, downloadedTarget)
	if util.HandleError(err) {
		return "", err
	}

	return processedTarget, nil
}

func (self *ManifestService) postprocessTarget(process string, target string) (filePath string, err error) {
	filePath = target
	return
}

func (self *ManifestService) packageForDeployment(packageType manifest.PackageType, location string, files []string) (output string, err error) {
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
		dest, err := util.MoveFile(archive, location)
		util.HandleError(err)
		output = "Package saved to " + dest
		fmt.Println(output)
	} else {
		fmt.Println("archive did not exist")
	}

	if stringUtils.IsBlank(output) {
		output = err.Error()
	}

	return output, nil
}

func (self *ManifestService) cleanup() {
	util.CleanUp()
}
