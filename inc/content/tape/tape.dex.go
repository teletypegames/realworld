package tape

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = TapeManager.Register(Tape{
	Name: inc.Dex,
	Label: inkwell.Text{
		{Lang: "en", Text: "DEX"},
		{Lang: "hu", Text: "DEX"},
	},
	Color:  theme.Amber,
	Dialog: inc.DialogDex,
})
