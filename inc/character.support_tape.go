package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	CharacterManager.Register(Character{
		Name:        TapeSupport,
		Label:       "TECH SUPPORT",
		Voice:       inkwell.VoiceLog,
		SpeechColor: RGB(0xE8C86A),
	})
}
