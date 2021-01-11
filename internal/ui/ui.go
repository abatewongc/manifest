package ui

import (
	gui "github.com/AllenDang/giu"
	"github.com/agrison/go-commons-lang/stringUtils"
)

var (
	targetedManifest string
)

func Start() {
	window := gui.NewMasterWindow("manifest", 640, 480, gui.MasterWindowFlagsNotResizable, nil)
	window.Run(loop)
}

func loop() {
	gui.SingleWindow("manifest").Layout(gui.Layout{
		gui.Line(gui.Button("Load Manifest").Size(620, 50).OnClick(onLoadManifest)),
		gui.Label("Loaded Manifest: " + stringUtils.Abbreviate(targetedManifest, 120)),
		gui.Line(gui.Box),

	})
}

func onLoadManifest() {
	targetedManifest = selectInputPath()
}