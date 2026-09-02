package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = CharacterManager.Register(Character{
	Name: inc.ThinFigure,
	Label: inkwell.Text{
		{Lang: "en", Text: "thin figure"},
		{Lang: "hu", Text: "vékony alak"},
	},
	Speed:       80,
	W:           20,
	H:           60,
	SpeechColor: theme.InkDim,
})
