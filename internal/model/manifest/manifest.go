package manifest

import (
	"encoding/json"
	"errors"
	"github.com/agrison/go-commons-lang/stringUtils"
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
	Location string      `json:"location"`
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
	rawJSON, err := util.ReadBytes(Path)
	util.HandleError(err)

	err = json.Unmarshal(rawJSON, &manifest)
	return
}

func (manifest *Manifest) Validate() (err error) {
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
