package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = CharacterManager.Register(Character{
	Name: inc.HotDogVendor,
	Label: inkwell.Text{
		{Lang: "en", Text: "hot dog vendor"},
		{Lang: "hu", Text: "hotdogárus"},
	},
	Speed:       80,
	W:           30,
	H:           66,
	SpeechColor: theme.AmberPale,
})
