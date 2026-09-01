package tape

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = TapeManager.Register(Tape{
	Name: inc.TapeSupport,
	Label: inkwell.Text{
		{Lang: "en", Text: "TECH SUPPORT"},
		{Lang: "hu", Text: "ÜGYFÉLSZOLGÁLAT"},
	},
	Color: theme.AmberPale,
})
