package web

import (
	"errors"
	"github.com/agrison/go-commons-lang/stringUtils"
	"github.com/aleosiss/manifest/internal/resource"
	"github.com/cavaliercoder/grab"
	"log"
)

var (
	ErrorCodes = [...]string{"403", "404"}
)


// DownloadTarget
func DownloadTarget(url string) (path string, err error) {
	localDest := resource.ManifestDownloadDir

	resp, err := grab.Get(localDest, url)
	err = handleError(err, url)

	path = resp.Filename
	if path != "" {
		log.Println("Download saved to " + path)
	}

	return
}

func handleError(err error, url string) error {
	if err == nil {
		return err
	}

	for _, code := range ErrorCodes {
		if stringUtils.Contains(err.Error(), code) {
			err = errors.New(err.Error() + " trying to reach " + url)
		}
	}

	return err
}