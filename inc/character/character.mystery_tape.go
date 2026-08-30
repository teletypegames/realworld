package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
	"git.teletypegames.org/games/realworld/inc/theme"
)

func init() {
	Manager.Register(Character{
		Name:        inc.TapeMystery,
		Label:       "UNKNOWN TAPE",
		Voice:       inkwell.VoiceLog,
		SpeechColor: theme.RGB(0xB9A98A),
	})
}
