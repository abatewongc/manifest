package web

import (
	"fmt"
	"github.com/aleosiss/manifest/internal/resource"
	"github.com/aleosiss/manifest/internal/util"
	"github.com/cavaliercoder/grab"
	"log"
)

// DownloadTarget
func DownloadTarget(url string) (path string, err error) {
	localDest := resource.ManifestDownloadDir

	resp, err := grab.Get(localDest, url)
	util.HandleError(err)

	path = resp.Filename

	log.Println(fmt.Sprintln("Download saved to", path))

	return
}