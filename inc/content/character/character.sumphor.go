package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
)

var Sumphor = CharacterManager.Register(&Character{
	Name: "sumphor",
	Label: inkwell.Text{
		{Lang: "en", Text: "Sumphor"},
		{Lang: "hu", Text: "Sumphor"},
	},
	Speed:       80,
	W:           26,
	H:           66,
	SpeechColor: theme.Amber,
})
