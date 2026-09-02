package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = CharacterManager.Register(Character{
	Name: inc.Secretary,
	Label: inkwell.Text{
		{Lang: "en", Text: "secretary"},
		{Lang: "hu", Text: "titkár"},
	},
	Speed:       80,
	W:           26,
	H:           66,
	SpeechColor: theme.Bone,
})
