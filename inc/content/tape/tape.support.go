package tape

import (
	"realworld/inc"
)

func init() {
	Manager.Register(Tape{
		Name: inc.TapeSupport,
	})
}
