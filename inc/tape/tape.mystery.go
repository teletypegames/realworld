package tape

import (
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Tape{
		Name:     inc.TapeMystery,
		Item:     inc.ItemMysteryTape,
		Dialogue: inc.DlgMysteryTape,
	})
}
