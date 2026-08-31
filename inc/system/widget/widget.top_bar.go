package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	Manager.Register(&inkwell.TopBar{
		Name:    "topbar",
		When:    whenPlaying,
		Height:  TopBarH,
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
