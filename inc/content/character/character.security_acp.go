package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
)

var SecurityACP = CharacterManager.Register(&Character{
	Name: "security_acp",
	Label: inkwell.Text{
		{Lang: "en", Text: "SECURITY ACP"},
		{Lang: "hu", Text: "BIZTONSÁGI ACP"},
	},
	Voice:       inkwell.VoiceLog,
	SpeechColor: theme.Green,
})
