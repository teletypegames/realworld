package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(&inkwell.TopBar{
		Name:    "topbar",
		When:    whenPlaying,
		Height:  TopBarH,
		TimeVar: inc.VarArmillaStrip,
		NoteVar: inc.VarNote,
	})
}
