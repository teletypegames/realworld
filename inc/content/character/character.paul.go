package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = CharacterManager.Register(Character{
	Name:  inc.Paul,
	Speed: 96,
	W:     28,
	H:     68,
	Start: inkwell.Point{
		X: 320,
		Y: 232,
	},
	SpeechColor: theme.Ink,
})
