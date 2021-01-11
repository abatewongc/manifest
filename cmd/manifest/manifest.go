package manifest

import (
	"encoding/json"
	"github.com/aleosiss/manifest/cmd/file"
	"github.com/aleosiss/manifest/internal/util"
)

// Manifest : This is why we're here.
type Manifest struct {
	Name    string   `json:"name"`
	Targets []Target `json:"targets"`
	Package Package  `json:"package"`
}

// Target : contains the friendly name of the artifact we're trying to
// download along with all of the metadata required to successfully download it.
type Target struct {
	Name          string `json:"name"`
	URL           string `json:"url"`
	TargetVersion string `json:"target_version"`
	PostProcess   string `json:"postprocess"`
}

type Package struct {
	Location string `json:"location"`
	Type     PackageType `json:"type"`
}

type PackageType string
const (
	ZIP = PackageType("zip")
)
var packageTypes = map[string]PackageType{
	"zip": ZIP,
	"ZIP": ZIP,
	"Zip": ZIP,
}

// From : deserialize cmd given file path
func From(Path string) (manifest Manifest, err error) {
	rawJSON, err := file.ReadBytes(Path)
	util.HandleError(err)

	err = json.Unmarshal(rawJSON, &manifest)
	return
}
