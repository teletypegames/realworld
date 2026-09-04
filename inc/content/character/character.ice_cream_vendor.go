package character

import (
	"realworld/inc/system/theme"
)

var IceCreamVendor = CharacterManager.Register(&Character{
	Name: "ice_cream_vendor",
	Label: Text{
		{Lang: "en", Text: "ice cream vendor"},
		{Lang: "hu", Text: "fagyiárus"},
	},
	Speed:       80,
	W:           26,
	H:           66,
	SpeechColor: theme.AmberPale,
})
