package tape

import (
	"realworld/inc"
)

func init() {
	TapeManager.Register(Tape{
		Name:   inc.TapeMystery,
		Item:   inc.ItemMysteryTape,
		Dialog: inc.DialogMysteryTape,
	})
}
