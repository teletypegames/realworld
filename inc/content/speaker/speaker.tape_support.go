package speaker

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
)

var TapeSupport = SpeakerManager.Register(&Speaker{
	Name: "tape_support",
	Label: Text{
		{Lang: "en", Text: "TECH SUPPORT"},
		{Lang: "hu", Text: "ÜGYFÉLSZOLGÁLAT"},
	},
	Delivery: inkwell.DeliveryLog,
	Color:    theme.AmberPale,
})
