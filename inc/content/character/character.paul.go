package character

import (
	"realworld/inc/system/theme"
)

var Paul = CharacterManager.Register(&Character{
	Name: "paul",
	Label: Text{
		{Lang: "en", Text: "Paul"},
		{Lang: "hu", Text: "Paul"},
	},
	Lore: Text{
		{Lang: "en", Text: "Fixes and sells other people's yesterdays: modems, a rotary phone, half a synth. Came to San Francisco because Noodle wrote instead of calling, and nobody sends you paper."},
		{Lang: "hu", Text: "Mások tegnapjait javítja és adja el: modemeket, egy tárcsás telefont, fél szintetizátort. Azért jött San Franciscóba, mert Noodle írt telefon helyett, és neked senki nem küld papírt."},
	},
	Speed: 96,
	W:     28,
	H:     68,
	Start: Point{
		X: 320,
		Y: 232,
	},
	SpeechColor: theme.Ink,
})
