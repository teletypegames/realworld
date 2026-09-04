package speaker

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
)

var Dex = SpeakerManager.Register(&Speaker{
	Name: "dex",
	Label: Text{
		{Lang: "en", Text: "DEX"},
		{Lang: "hu", Text: "DEX"},
	},
	Voice: inkwell.VoiceLog,
	Color: theme.Amber,
})
