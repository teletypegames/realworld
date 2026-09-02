package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = CharacterManager.Register(Character{
	Name: inc.Receptionist,
	Label: inkwell.Text{
		{Lang: "en", Text: "receptionist"},
		{Lang: "hu", Text: "recepciós"},
	},
	Speed:       80,
	W:           26,
	H:           60,
	SpeechColor: theme.Bone,
})
