package resource

import "os"

func CreateDirectories() {
	_ = os.Mkdir(ManifestAppDir, os.ModeDir)
	_ = os.Mkdir(ManifestDownloadDir, os.ModeDir)
	_ = os.Mkdir(ManifestStagingDir, os.ModeDir)
}
