package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
)

var Noodle = CharacterManager.Register(&Character{
	Name: "noodle",
	Label: inkwell.Text{
		{Lang: "en", Text: "Noodle"},
		{Lang: "hu", Text: "Noodle"},
	},
	Speed:       80,
	W:           40,
	H:           62,
	SpeechColor: theme.Bone,
})
