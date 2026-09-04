package character

import (
	"realworld/inc/system/theme"
)

var ThinFigure = CharacterManager.Register(&Character{
	Name: "thin_figure",
	Label: Text{
		{Lang: "en", Text: "thin figure"},
		{Lang: "hu", Text: "vékony alak"},
	},
	Speed:       80,
	W:           20,
	H:           60,
	SpeechColor: theme.InkDim,
})
