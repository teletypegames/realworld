package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
	"git.teletypegames.org/games/realworld/inc/theme"
)

func init() {
	Manager.Register(Character{
		Name:        inc.Dex,
		Label:       "DEX",
		Voice:       inkwell.VoiceLog,
		SpeechColor: theme.Amber,
	})
}
