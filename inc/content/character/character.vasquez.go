package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = CharacterManager.Register(Character{
	Name: inc.Vasquez,
	Label: inkwell.Text{
		{Lang: "en", Text: "Det. Vasquez"},
		{Lang: "hu", Text: "Vasquez nyomozó"},
	},
	Speed:       80,
	W:           26,
	H:           68,
	SpeechColor: theme.Amber,
})
