package character

import (
	"realworld/inc/system/theme"
)

var Receptionist = CharacterManager.Register(&Character{
	Name: "receptionist",
	Label: Text{
		{Lang: "en", Text: "receptionist"},
		{Lang: "hu", Text: "recepciós"},
	},
	Speed:       80,
	W:           26,
	H:           60,
	SpeechColor: theme.Bone,
})
