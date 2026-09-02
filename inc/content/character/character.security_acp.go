package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = CharacterManager.Register(Character{
	Name: inc.SecurityACP,
	Label: inkwell.Text{
		{Lang: "en", Text: "SECURITY ACP"},
		{Lang: "hu", Text: "BIZTONSÁGI ACP"},
	},
	Voice:       inkwell.VoiceLog,
	SpeechColor: theme.Green,
})
