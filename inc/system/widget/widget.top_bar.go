package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	WidgetManager.Register(&inkwell.TopBar{
		Name:    "topbar",
		When:    whenPlaying,
		Height:  TopBarHeight,
		TimeVar: inc.VarArmillaStrip,
		NoteVar: inc.VarNote,
		Notes: map[string]inkwell.Text{
			inc.NotePaused: {
				{Lang: "en", Text: "paused"},
				{Lang: "hu", Text: "szünet"},
			},
		},
	})
}
