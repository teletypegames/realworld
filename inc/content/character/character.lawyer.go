package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
)

var Lawyer = CharacterManager.Register(&Character{
	Name: "lawyer",
	Label: inkwell.Text{
		{Lang: "en", Text: "lawyer"},
		{Lang: "hu", Text: "ügyvéd"},
	},
	Speed:       80,
	W:           26,
	H:           62,
	SpeechColor: theme.InkDim,
})
