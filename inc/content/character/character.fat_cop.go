package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
)

var FatCop = CharacterManager.Register(&Character{
	Name: "fat_cop",
	Label: inkwell.Text{
		{Lang: "en", Text: "fat cop"},
		{Lang: "hu", Text: "kövér rendőr"},
	},
	Speed:       80,
	W:           40,
	H:           64,
	SpeechColor: theme.InkDim,
})
