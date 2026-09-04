package character

import (
	"realworld/inc/system/theme"
)

var Lawyer = CharacterManager.Register(&Character{
	Name: "lawyer",
	Label: Text{
		{Lang: "en", Text: "lawyer"},
		{Lang: "hu", Text: "ügyvéd"},
	},
	Speed:       80,
	W:           26,
	H:           62,
	SpeechColor: theme.InkDim,
})
