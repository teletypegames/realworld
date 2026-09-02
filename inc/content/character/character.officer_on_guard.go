package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = CharacterManager.Register(Character{
	Name: inc.OfficerOnGuard,
	Label: inkwell.Text{
		{Lang: "en", Text: "officer"},
		{Lang: "hu", Text: "rendőr"},
	},
	Speed:       80,
	W:           34,
	H:           70,
	SpeechColor: theme.InkDim,
})
