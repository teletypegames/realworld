package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"

	"realworld/inc/system/widget"
)

var Game = StageManager.Register(&Stage{
	Name:  "game",
	World: true,
	Title: Text{
		{Lang: "en", Text: "REAL WORLD"},
		{Lang: "hu", Text: "REAL WORLD"},
	},
})

var _ = widget.WidgetManager.Register(&inkwell.KeyBinding{
	Name:         "pause",
	When:         inkwell.OnStage(Game),
	Key:          ebiten.KeyEscape,
	Do:           inkwell.GoToStage(MainMenu),
	DuringScript: true,
})

var _ = widget.WidgetManager.Register(&inkwell.KeyBinding{
	Name:         "dossier_open",
	When:         inkwell.OnStage(Game),
	Key:          ebiten.KeyTab,
	Do:           inkwell.GoToStage(Dossier),
	DuringScript: true,
})
