package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"

	"realworld/inc"
	"realworld/inc/system/widget"
)

var _ = StageManager.Register(Stage{
	Name:  inc.StageGame,
	World: true,
	Title: inkwell.Text{
		{Lang: "en", Text: "REAL WORLD"},
		{Lang: "hu", Text: "REAL WORLD"},
	},
})

var _ = widget.WidgetManager.Register(&inkwell.KeyBinding{
	Name:         "pause",
	When:         inkwell.OnStage(inc.StageGame),
	Key:          ebiten.KeyEscape,
	Do:           inkwell.GoToStage(inc.StageMainMenu),
	DuringScript: true,
})
