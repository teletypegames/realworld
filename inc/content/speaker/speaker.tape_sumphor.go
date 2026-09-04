package speaker

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
)

var TapeSumphor = SpeakerManager.Register(&Speaker{
	Name: "tape_sumphor",
	Label: inkwell.Text{
		{Lang: "en", Text: "SUMPHOR'S TAPE"},
		{Lang: "hu", Text: "SUMPHOR KAZETTÁJA"},
	},
	Voice: inkwell.VoiceLog,
	Color: theme.InkDim,
})
