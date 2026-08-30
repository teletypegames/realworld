package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	WidgetManager.Register(&inkwell.TopBar{
		Name:    "topbar",
		When:    whenPlaying,
		Height:  TopBarH,
		TimeVar: VarArmillaStrip,
		NoteVar: VarNote,
	})
}
