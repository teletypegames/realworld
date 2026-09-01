package tape

import (
	"realworld/inc"
)

var _ = TapeManager.Register(Tape{
	Name: inc.TapeSupport,
})
