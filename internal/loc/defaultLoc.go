package loc

var LoadingManifest = "Loaded Manifest: "
var NoFileErrors = "No file errors reported."

func LocalizeFileError(err error) string {
	return err.Error()
}