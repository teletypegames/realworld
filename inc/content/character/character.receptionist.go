package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
)

var Receptionist = CharacterManager.Register(&Character{
	Name: "receptionist",
	Label: inkwell.Text{
		{Lang: "en", Text: "receptionist"},
		{Lang: "hu", Text: "recepciós"},
	},
	Speed:       80,
	W:           26,
	H:           60,
	SpeechColor: theme.Bone,
})
