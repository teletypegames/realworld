package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/util"
)

func init() {
	CharacterManager.Register(Character{
		Name: inc.TapeSupport,
		Label: inkwell.Text{
			{Lang: "en", Text: "TECH SUPPORT"},
			{Lang: "hu", Text: "ÜGYFÉLSZOLGÁLAT"},
		},
		Voice:       inkwell.VoiceLog,
		SpeechColor: util.RGB(0xE8C86A),
	})
}
