package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = CharacterManager.Register(Character{
	Name: inc.Norman,
	Label: inkwell.Text{
		{Lang: "en", Text: "Norman"},
		{Lang: "hu", Text: "Norman"},
	},
	Speed:       80,
	W:           26,
	H:           66,
	SpeechColor: theme.AmberPale,
})
