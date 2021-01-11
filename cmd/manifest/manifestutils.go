package manifest

import (
	"errors"
	"fmt"
	"github.com/agrison/go-commons-lang/stringUtils"
	"github.com/aleosiss/manifest/internal/util"
	"github.com/aleosiss/manifest/internal/web"
	"os"
	"strings"
	"sync"
)

func Validate(manifest Manifest) (err error) {
	// nothing too complicated
	if stringUtils.IsBlank(manifest.Name)  {
		err = errors.New("manifest requires a name to be considered valid")
		return
	}

	if stringUtils.IsBlank(manifest.Package.Location) {
		err = errors.New("manifest needs a package location to put the packaged manifest in")
		return
	}

	if  _, ok := packageTypes[string(manifest.Package.Type)]; !ok {
		err = errors.New("manifest does not have a valid PackageType")
	}

	if len(manifest.Targets) < 1 {
		err = errors.New("manifest needs at least one target")
	}

	return nil
}

func Manifesto(filePath string) (err error, fileErrors []error) {
	manifestFile, err := From(filePath)
	if util.HandleError(err) {
		return err, fileErrors
	}

	err = Validate(manifestFile)

	var files []string

	fmt.Println("Handling Manifest: " + manifestFile.Name)
	wg := sync.WaitGroup{}
	for _, target := range manifestFile.Targets {
		wg.Add(1)
		go func(target Target) {
			defer wg.Done()
			file, err := handleTarget(target)

			if err != nil { fileErrors = append(fileErrors, err) }
			if file != "" { files = append(files, file) }
		}(target)
	}
	wg.Wait()

	err = packageForDeployment(manifestFile.Package.Type, manifestFile.Package.Location, files)
	if util.HandleError(err) {
		return err, fileErrors
	}

	cleanup()
	return err, fileErrors
}

func handleTarget(target Target) (string, error) {
	fmt.Println("Found target: " + target.Name)
	url, err := util.ExpandText(target.URL, "version", target.TargetVersion)
	if util.HandleError(err) {
		return "", err
	}

	downloadedTarget, err := web.DownloadTarget(url)
	if util.HandleError(err) {
		return "", err
	}

	processedTarget, err := postprocessTarget(target.PostProcess, downloadedTarget)
	if util.HandleError(err) {
		return "", err
	}

	return processedTarget, nil
}

func postprocessTarget(process string, target string) (filePath string, err error) {
	filePath = target
	return
}

func packageForDeployment(packageType PackageType, location string, files []string) (err error) {
	var archive string

	if packageType == ZIP {
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

