package web

import (
	"fmt"
	"github.com/aleosiss/manifest/cmd/manifest"
	"github.com/aleosiss/manifest/internal/resource"
	"github.com/cavaliercoder/grab"
	"log"
)

// DownloadTarget
func DownloadTarget(target manifest.Target, url string) (path string, err error) {
	localDest := resource.ManifestDownloadDir

	resp, err := grab.Get(localDest, url)
	if err != nil {
		log.Fatal(err)
	}

	path = resp.Filename

	log.Println(fmt.Sprintln("Download saved to", path))

	return
}