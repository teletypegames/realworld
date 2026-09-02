package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = CharacterManager.Register(Character{
	Name: inc.IceCreamVendor,
	Label: inkwell.Text{
		{Lang: "en", Text: "ice cream vendor"},
		{Lang: "hu", Text: "fagyiárus"},
	},
	Speed:       80,
	W:           26,
	H:           66,
	SpeechColor: theme.AmberPale,
})
