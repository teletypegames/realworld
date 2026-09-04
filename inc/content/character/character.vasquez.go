package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
)

var Vasquez = CharacterManager.Register(&Character{
	Name: "vasquez",
	Label: inkwell.Text{
		{Lang: "en", Text: "Det. Vasquez"},
		{Lang: "hu", Text: "Vasquez nyomozó"},
	},
	Speed:       80,
	W:           26,
	H:           68,
	SpeechColor: theme.Amber,
})
