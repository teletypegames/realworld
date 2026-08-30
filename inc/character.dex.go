package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	CharacterManager.Register(Character{
		Name:        Dex,
		Label:       "DEX",
		Voice:       inkwell.VoiceLog,
		SpeechColor: Amber,
	})
}
