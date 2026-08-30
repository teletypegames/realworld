package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	CharacterManager.Register(Character{
		Name:        TapeMystery,
		Label:       "UNKNOWN TAPE",
		Voice:       inkwell.VoiceLog,
		SpeechColor: RGB(0xB9A98A),
	})
}
