package tape

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = TapeManager.Register(Tape{
	Name: inc.TapeMystery,
	Item: inc.ItemMysteryTape,
	Label: inkwell.Text{
		{Lang: "en", Text: "UNKNOWN TAPE"},
		{Lang: "hu", Text: "ISMERETLEN KAZETTA"},
	},
	Color:  theme.Bone,
	Dialog: inc.DialogMysteryTape,
})
