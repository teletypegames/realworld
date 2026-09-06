package character

import (
	"realworld/inc/system/theme"
)

var Norman = CharacterManager.Register(&Character{
	Name: "norman",
	Label: Text{
		{Lang: "en", Text: "Norman"},
		{Lang: "hu", Text: "Norman"},
	},
	Lore: Text{
		{Lang: "en", Text: "Norman Reds. A developer on the BVK's human simulation, the one name on the list with no whereabouts. Never believed he belonged in the room; his own tape, he wrote, was the clever one."},
		{Lang: "hu", Text: "Norman Reds. A BVK humán szimulációjának fejlesztője, az egyetlen név a listán tartózkodási hely nélkül. Sosem hitte, hogy joga van ott lenni; a saját kazettája, írta, az okosabb."},
	},
	Speed:       80,
	W:           26,
	H:           66,
	SpeechColor: theme.AmberPale,
})
