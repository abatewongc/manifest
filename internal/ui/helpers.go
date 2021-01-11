package ui

import "github.com/sqweek/dialog"

func selectInputPath() string {
	path, _ := dialog.File().Filter("manifest", "json").Load()
	return path
}
