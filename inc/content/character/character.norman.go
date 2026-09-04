package character

import (
	"realworld/inc/system/theme"
)

var Norman = CharacterManager.Register(&Character{
	Name: "norman",
	Label: Text{
		{Lang: "en", Text: "Norman"},
		{Lang: "hu", Text: "Norman"},
	},
	Speed:       80,
	W:           26,
	H:           66,
	SpeechColor: theme.AmberPale,
})
