package tape

import (
	"realworld/inc"
)

func init() {
	Manager.Register(Tape{
		Name:     inc.TapeMystery,
		Item:     inc.ItemMysteryTape,
		Dialogue: inc.DlgMysteryTape,
	})
}
