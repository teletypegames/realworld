package tape

import (
	"realworld/inc"
)

var _ = TapeManager.Register(Tape{
	Name:   inc.TapeMystery,
	Item:   inc.ItemMysteryTape,
	Dialog: inc.DialogMysteryTape,
})
