package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	CharacterManager.Register(Character{
		Name:  Paul,
		Speed: 96,
		W:     28,
		H:     68,
		Start: inkwell.Point{
			X: 120,
			Y: 232,
		},
		SpeechColor: Ink,
	})
}
