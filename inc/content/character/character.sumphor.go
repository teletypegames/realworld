package character

import (
	"realworld/inc/system/theme"
)

var Sumphor = CharacterManager.Register(&Character{
	Name: "sumphor",
	Label: Text{
		{Lang: "en", Text: "Sumphor"},
		{Lang: "hu", Text: "Sumphor"},
	},
	Speed:       80,
	W:           26,
	H:           66,
	SpeechColor: theme.Amber,
})
