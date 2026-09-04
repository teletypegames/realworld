package speaker

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
)

var TapeServerFarm = SpeakerManager.Register(&Speaker{
	Name: "tape_server_farm",
	Label: Text{
		{Lang: "en", Text: "ACCESS TAPE"},
		{Lang: "hu", Text: "BELÉPŐKAZETTA"},
	},
	Voice: inkwell.VoiceLog,
	Color: theme.AmberLow,
})
