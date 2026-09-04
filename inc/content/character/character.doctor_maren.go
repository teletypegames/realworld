package character

import (
	"realworld/inc/system/theme"
)

var DoctorMaren = CharacterManager.Register(&Character{
	Name: "doctor_maren",
	Label: Text{
		{Lang: "en", Text: "Dr. Maren"},
		{Lang: "hu", Text: "Dr. Maren"},
	},
	Speed:       80,
	W:           26,
	H:           68,
	SpeechColor: theme.Bone,
})
