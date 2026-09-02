package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = CharacterManager.Register(Character{
	Name: inc.Shopkeeper,
	Label: inkwell.Text{
		{Lang: "en", Text: "shopkeeper"},
		{Lang: "hu", Text: "boltos"},
	},
	Speed:       80,
	W:           26,
	H:           62,
	SpeechColor: theme.Amber,
})
