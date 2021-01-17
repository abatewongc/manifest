package ui

import (
	gui "github.com/AllenDang/giu"
	"github.com/agrison/go-commons-lang/stringUtils"
	"github.com/aleosiss/manifest/internal/globals"
	"github.com/aleosiss/manifest/internal/loc"
	"github.com/aleosiss/manifest/internal/service"
)


type UI struct {
	targetedManifest string
	manifestService service.ManifestService
}

func New(_manifestService *service.ManifestService) UI {
	var g UI
	g.manifestService = *_manifestService
	return g
}

func (self *UI) Start() {
	window := gui.NewMasterWindow("manifest", 640, 480, gui.MasterWindowFlagsNotResizable, nil)
	window.Run(self.loop)
}

func (self *UI) loop() {
	gui.SingleWindow("manifest").Layout(gui.Layout{
		gui.Line(gui.Button("Load Manifest").Size(620, 50).OnClick(self.onLoadManifest)),
		gui.Label(loc.LoadingManifest + stringUtils.Abbreviate(self.targetedManifest, 120)),
		gui.Line(gui.Condition(globals.UIWorking, gui.Layout{
			gui.ProgressIndicator("isWorking", globals.UIProgressBarLabel, 620, 390, 120),
		}, gui.Layout{
			gui.Label(globals.UIMsgBoxContents).Wrapped(true),
		})),
	})
}


func (self *UI) onLoadManifest() {
	if globals.UIWorking {
		return
	}

	self.targetedManifest = selectInputPath()
	go self.handleManifest()
}

func (self *UI) handleManifest() {
	globals.UIWorking = true
	gui.Update()
	_, _ = self.manifestService.Process(self.targetedManifest, true)
	gui.Update()
	globals.UIWorking = false
	gui.Update()
	self.targetedManifest = ""
	gui.Update()
}