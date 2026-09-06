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
	Lore: Text{
		{Lang: "en", Text: "Whoever sits in the passenger seat of the cruiser. Keeps the fat cop talking, and never gets out."},
		{Lang: "hu", Text: "Akárki is ül a járőrkocsi anyósülésén. Beszélteti a kövér rendőrt, és sosem száll ki."},
	},
	Speed:       80,
	W:           20,
	H:           60,
	SpeechColor: theme.InkDim,
})
