package tape

import (
	"realworld/inc"
)

func init() {
	TapeManager.Register(Tape{
		Name:     inc.Dex,
		Dialogue: inc.DialogDex,
	})
}
