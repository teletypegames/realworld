package character

import (
	"realworld/inc/system/theme"
)

var FatCop = CharacterManager.Register(&Character{
	Name: "fat_cop",
	Label: Text{
		{Lang: "en", Text: "fat cop"},
		{Lang: "hu", Text: "kövér rendőr"},
	},
	Speed:       80,
	W:           40,
	H:           64,
	SpeechColor: theme.InkDim,
})
