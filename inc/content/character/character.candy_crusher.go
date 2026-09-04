package character

import (
	"realworld/inc/system/theme"
)

var CandyCrusher = CharacterManager.Register(&Character{
	Name: "candy_crusher",
	Label: Text{
		{Lang: "en", Text: "man with the handheld"},
		{Lang: "hu", Text: "kézikonzolos fickó"},
	},
	Speed:       80,
	W:           30,
	H:           60,
	SpeechColor: theme.Bone,
})
