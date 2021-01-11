package manifest

import (
	"errors"
	"github.com/agrison/go-commons-lang/stringUtils"
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