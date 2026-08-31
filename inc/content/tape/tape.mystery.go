package tape

import (
	"realworld/inc"
)

func init() {
	TapeManager.Register(Tape{
		Name:     inc.TapeMystery,
		Item:     inc.ItemMysteryTape,
		Dialogue: inc.DlgMysteryTape,
	})
}
