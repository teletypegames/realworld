package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

func init() {
	Manager.Register(Character{
		Name: inc.Dex,
		Label: inkwell.Text{
			{Lang: "en", Text: "DEX"},
			{Lang: "hu", Text: "DEX"},
		},
		Voice:       inkwell.VoiceLog,
		SpeechColor: theme.Amber,
	})
}
