package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = CharacterManager.Register(Character{
	Name: inc.Noodle,
	Label: inkwell.Text{
		{Lang: "en", Text: "Noodle"},
		{Lang: "hu", Text: "Noodle"},
	},
	Speed:       80,
	W:           40,
	H:           62,
	SpeechColor: theme.Bone,
})
