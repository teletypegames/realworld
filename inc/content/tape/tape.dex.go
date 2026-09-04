package tape

import (
	"realworld/inc/content/dialog"
	"realworld/inc/content/speaker"
)

var Dex = TapeManager.Register(&Tape{
	Name:    "dex",
	Speaker: speaker.Dex,
	Dialog:  dialog.Dex,
})
