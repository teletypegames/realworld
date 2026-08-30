package tape

import (
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Tape{
		Name:     inc.Dex,
		Dialogue: inc.DlgDex,
	})
}
