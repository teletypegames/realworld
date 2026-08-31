package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

func init() {
	CharacterManager.Register(Character{
		Name: inc.TapeSupport,
		Label: inkwell.Text{
			{Lang: "en", Text: "TECH SUPPORT"},
			{Lang: "hu", Text: "ÜGYFÉLSZOLGÁLAT"},
		},
		Voice:       inkwell.VoiceLog,
		SpeechColor: theme.RGB(0xE8C86A),
	})
}
