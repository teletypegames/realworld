package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
)

var IceCreamVendor = CharacterManager.Register(&Character{
	Name: "ice_cream_vendor",
	Label: inkwell.Text{
		{Lang: "en", Text: "ice cream vendor"},
		{Lang: "hu", Text: "fagyiárus"},
	},
	Speed:       80,
	W:           26,
	H:           66,
	SpeechColor: theme.AmberPale,
})
