package speaker

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
)

var TapeCommentary = SpeakerManager.Register(&Speaker{
	Name: "tape_commentary",
	Label: inkwell.Text{
		{Lang: "en", Text: "DEV COMMENTARY"},
		{Lang: "hu", Text: "FEJLESZTŐI KOMMENTÁR"},
	},
	Voice: inkwell.VoiceLog,
	Color: theme.Amber,
})
