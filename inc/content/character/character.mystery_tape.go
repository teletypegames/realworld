package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/util"
)

func init() {
	CharacterManager.Register(Character{
		Name: inc.TapeMystery,
		Label: inkwell.Text{
			{Lang: "en", Text: "UNKNOWN TAPE"},
			{Lang: "hu", Text: "ISMERETLEN KAZETTA"},
		},
		Voice:       inkwell.VoiceLog,
		SpeechColor: util.RGB(0xB9A98A),
	})
}
