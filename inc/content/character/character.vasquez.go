package character

import (
	"realworld/inc/system/theme"
)

var Vasquez = CharacterManager.Register(&Character{
	Name: "vasquez",
	Label: Text{
		{Lang: "en", Text: "Det. Vasquez"},
		{Lang: "hu", Text: "Vasquez nyomozó"},
	},
	Speed:       80,
	W:           26,
	H:           68,
	SpeechColor: theme.Amber,
})
