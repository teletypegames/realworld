package character

import (
	"realworld/inc/system/theme"
)

var WorriedRelative = CharacterManager.Register(&Character{
	Name: "worried_relative",
	Label: Text{
		{Lang: "en", Text: "worried woman"},
		{Lang: "hu", Text: "aggódó nő"},
	},
	Speed:       80,
	W:           26,
	H:           62,
	SpeechColor: theme.Bone,
})
