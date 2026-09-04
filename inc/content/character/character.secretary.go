package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
)

var Secretary = CharacterManager.Register(&Character{
	Name: "secretary",
	Label: inkwell.Text{
		{Lang: "en", Text: "secretary"},
		{Lang: "hu", Text: "titkár"},
	},
	Speed:       80,
	W:           26,
	H:           66,
	SpeechColor: theme.Bone,
})
