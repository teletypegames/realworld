package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = CharacterManager.Register(Character{
	Name: inc.CandyCrusher,
	Label: inkwell.Text{
		{Lang: "en", Text: "man with the handheld"},
		{Lang: "hu", Text: "kézikonzolos fickó"},
	},
	Speed:       80,
	W:           30,
	H:           60,
	SpeechColor: theme.Bone,
})
