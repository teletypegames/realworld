package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
)

var Paul = CharacterManager.Register(&Character{
	Name:  "paul",
	Speed: 96,
	W:     28,
	H:     68,
	Start: inkwell.Point{
		X: 320,
		Y: 232,
	},
	SpeechColor: theme.Ink,
})
