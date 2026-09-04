package character

import (
	"realworld/inc/system/theme"
)

var OfficerOnGuard = CharacterManager.Register(&Character{
	Name: "officer_on_guard",
	Label: Text{
		{Lang: "en", Text: "officer"},
		{Lang: "hu", Text: "rendőr"},
	},
	Speed:       80,
	W:           34,
	H:           70,
	SpeechColor: theme.InkDim,
})
