package util

import (
	"fmt"
	"github.com/aleosiss/manifest/internal/resource"
	"github.com/draxil/gomv"
	"os"
	"path/filepath"
)

func CleanUp() {
	err := os.Remove(resource.ManifestDownloadDir)
	_ = os.Mkdir(resource.ManifestDownloadDir, os.ModeDir)

	err = os.Remove(resource.ManifestStagingDir)
	_ = os.Mkdir(resource.ManifestStagingDir, os.ModeDir)

	if err != nil {
		fmt.Println(err)
	}
}

func Exists(name string) (b bool) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	b = err == nil
	return b
}

func MoveFile(src string, dest string) (err error) {
	srcfi, err := os.Stat(src)
	destfi, err := os.Stat(dest)

	if err != nil {
		return
	}

	var filename string
	if ! srcfi.IsDir() && destfi.IsDir() {
		filename = filepath.Base(src)
	}

	dest = filepath.Join(dest, filename)

	err = gomv.MoveFile(src, dest)
	return
}
