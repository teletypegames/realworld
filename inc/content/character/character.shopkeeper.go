package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
)

var Shopkeeper = CharacterManager.Register(&Character{
	Name: "shopkeeper",
	Label: inkwell.Text{
		{Lang: "en", Text: "shopkeeper"},
		{Lang: "hu", Text: "boltos"},
	},
	Speed:       80,
	W:           26,
	H:           62,
	SpeechColor: theme.Amber,
})
