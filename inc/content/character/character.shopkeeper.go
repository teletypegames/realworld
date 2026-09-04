package character

import (
	"realworld/inc/system/theme"
)

var Shopkeeper = CharacterManager.Register(&Character{
	Name: "shopkeeper",
	Label: Text{
		{Lang: "en", Text: "shopkeeper"},
		{Lang: "hu", Text: "boltos"},
	},
	Speed:       80,
	W:           26,
	H:           62,
	SpeechColor: theme.Amber,
})
