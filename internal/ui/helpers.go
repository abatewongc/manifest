package ui

import (
	"github.com/aleosiss/manifest/internal/globals"
	"github.com/sqweek/dialog"
	"os"
)

func selectInputPath() string {
	path, _ := dialog.File().Filter("manifest", "json").Load()

	// using dialog changes our cwd, so reset it
	_ = os.Chdir(globals.CWD)

	return path
}
