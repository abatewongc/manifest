package util

import (
	"github.com/aleosiss/manifest/internal/resource"
	"github.com/draxil/gomv"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CleanUp() {
	err := os.RemoveAll(resource.ManifestDownloadDir)
	_ = os.Mkdir(resource.ManifestDownloadDir, os.ModeDir)

	err = os.RemoveAll(resource.ManifestStagingDir)
	_ = os.Mkdir(resource.ManifestStagingDir, os.ModeDir)

	HandleError(err)
}

func Exists(name string) (b bool) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	b = err == nil
	return b
}

func MoveFile(src string, dest string) (_ string, err error) {
	srcfi, err := os.Stat(src)
	destfi, err := os.Stat(dest)

	if err != nil {
		return
	}

	var filename string
	if !srcfi.IsDir() && destfi.IsDir() {
		filename = filepath.Base(src)
	}

	dest = filepath.Join(dest, filename)

	err = gomv.MoveFile(src, dest)
	dest, _ = filepath.Abs(dest)
	return dest, err
}

// ReadBytes : Beep beep!
func ReadBytes(filePath string) (data []byte, err error) {
	data, err = ioutil.ReadFile(filePath)
	return
}

