package resource

import (
	"github.com/vrischmann/userdir"
	"path/filepath"
)

var ManifestAppDir = filepath.FromSlash(userdir.GetDataHome()) + string (filepath.Separator) + "manifestapp"
var ManifestDownloadDir = ManifestAppDir + string (filepath.Separator) + "download"
var ManifestStagingDir = ManifestAppDir + string (filepath.Separator) + "staging"