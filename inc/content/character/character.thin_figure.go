package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
)

var ThinFigure = CharacterManager.Register(&Character{
	Name: "thin_figure",
	Label: inkwell.Text{
		{Lang: "en", Text: "thin figure"},
		{Lang: "hu", Text: "vékony alak"},
	},
	Speed:       80,
	W:           20,
	H:           60,
	SpeechColor: theme.InkDim,
})
