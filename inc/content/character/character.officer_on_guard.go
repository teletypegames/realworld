package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
)

var OfficerOnGuard = CharacterManager.Register(&Character{
	Name: "officer_on_guard",
	Label: inkwell.Text{
		{Lang: "en", Text: "officer"},
		{Lang: "hu", Text: "rendőr"},
	},
	Speed:       80,
	W:           34,
	H:           70,
	SpeechColor: theme.InkDim,
})
