package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = CharacterManager.Register(Character{
	Name: inc.FatCop,
	Label: inkwell.Text{
		{Lang: "en", Text: "fat cop"},
		{Lang: "hu", Text: "kövér rendőr"},
	},
	Speed:       80,
	W:           40,
	H:           64,
	SpeechColor: theme.InkDim,
})
