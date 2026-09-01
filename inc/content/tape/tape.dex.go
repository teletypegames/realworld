package tape

import (
	"realworld/inc"
)

var _ = TapeManager.Register(Tape{
	Name:   inc.Dex,
	Dialog: inc.DialogDex,
})
