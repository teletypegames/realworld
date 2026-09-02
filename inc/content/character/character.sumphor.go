package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = CharacterManager.Register(Character{
	Name: inc.Sumphor,
	Label: inkwell.Text{
		{Lang: "en", Text: "Sumphor"},
		{Lang: "hu", Text: "Sumphor"},
	},
	Speed:       80,
	W:           26,
	H:           66,
	SpeechColor: theme.Amber,
})
