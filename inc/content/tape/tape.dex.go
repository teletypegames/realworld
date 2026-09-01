package tape

import (
	"realworld/inc"
)

func init() {
	TapeManager.Register(Tape{
		Name:   inc.Dex,
		Dialog: inc.DialogDex,
	})
}
