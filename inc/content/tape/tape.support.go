package tape

import (
	"realworld/inc"
)

func init() {
	TapeManager.Register(Tape{
		Name: inc.TapeSupport,
	})
}
