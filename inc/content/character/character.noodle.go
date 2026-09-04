package character

import (
	"realworld/inc/system/theme"
)

var Noodle = CharacterManager.Register(&Character{
	Name: "noodle",
	Label: Text{
		{Lang: "en", Text: "Noodle"},
		{Lang: "hu", Text: "Noodle"},
	},
	Speed:       80,
	W:           40,
	H:           62,
	SpeechColor: theme.Bone,
})
