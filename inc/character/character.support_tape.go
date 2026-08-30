package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
	"git.teletypegames.org/games/realworld/inc/theme"
)

func init() {
	Manager.Register(Character{
		Name:        inc.TapeSupport,
		Label:       "TECH SUPPORT",
		Voice:       inkwell.VoiceLog,
		SpeechColor: theme.RGB(0xE8C86A),
	})
}
