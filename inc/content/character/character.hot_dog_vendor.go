package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
)

var HotDogVendor = CharacterManager.Register(&Character{
	Name: "hot_dog_vendor",
	Label: inkwell.Text{
		{Lang: "en", Text: "hot dog vendor"},
		{Lang: "hu", Text: "hotdogárus"},
	},
	Speed:       80,
	W:           30,
	H:           66,
	SpeechColor: theme.AmberPale,
})
