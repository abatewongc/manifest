package util

import (
	"compress/flate"
	"github.com/aleosiss/manifest/internal/resource"
	"github.com/mholt/archiver/v3"
	"path/filepath"
)

var zip = archiver.Zip {
	CompressionLevel:       flate.DefaultCompression,
	MkdirAll:               true,
	SelectiveCompression:   true,
	ContinueOnError:        false,
	OverwriteExisting:      false,
	ImplicitTopLevelFolder: false,
}



func ArchiveZip(files []string) (archive string, err error) {
	archive = resource.ManifestStagingDir + string(filepath.Separator) +  "manifest.zip"
	err = zip.Archive(files, archive)
	return
}
