package speaker

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
)

var TapeMystery = SpeakerManager.Register(&Speaker{
	Name: "tape_mystery",
	Label: Text{
		{Lang: "en", Text: "UNKNOWN TAPE"},
		{Lang: "hu", Text: "ISMERETLEN KAZETTA"},
	},
	Voice: inkwell.VoiceLog,
	Color: theme.Bone,
})
