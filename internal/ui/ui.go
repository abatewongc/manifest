package ui

import (
	gui "github.com/AllenDang/giu"
	"github.com/agrison/go-commons-lang/stringUtils"
	"github.com/aleosiss/manifest/cmd/manifest"
	"github.com/aleosiss/manifest/internal/globals"
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
		gui.Line(gui.Condition(globals.UIWorking, gui.Layout{
			gui.ProgressIndicator("isWorking", globals.UIProgressBarLabel, 620, 390, 120),
		}, nil)),
	})
}


func onLoadManifest() {
	if globals.UIWorking {
		return
	}

	targetedManifest = selectInputPath()
	go handleManifest()
}

func handleManifest() {
	globals.UIWorking = true
	gui.Update()
	_, _ = manifest.Manifesto(targetedManifest, true)
	gui.Update()
	globals.UIWorking = false
	gui.Update()
	targetedManifest = ""
}