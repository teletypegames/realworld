package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = CharacterManager.Register(Character{
	Name: inc.DoctorMaren,
	Label: inkwell.Text{
		{Lang: "en", Text: "Dr. Maren"},
		{Lang: "hu", Text: "Dr. Maren"},
	},
	Speed:       80,
	W:           26,
	H:           68,
	SpeechColor: theme.Bone,
})
