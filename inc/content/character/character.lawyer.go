package character

import (
	"realworld/inc/system/theme"
)

var Lawyer = CharacterManager.Register(&Character{
	Name: "lawyer",
	Label: Text{
		{Lang: "en", Text: "lawyer"},
		{Lang: "hu", Text: "ügyvéd"},
	},
	Lore: Text{
		{Lang: "en", Text: "Suit, briefcase, tired face, on the bench in the waiting room. Mutters about a Personal Tape somebody bought off the market and shouldn't have. Not now. Waiting."},
		{Lang: "hu", Text: "Öltöny, aktatáska, fáradt arc, a váróterem padján. Egy Personal Tape-ről motyog, amit valaki a piacon vett, és nem kellett volna. Most nem alkalmas. Várakozik."},
	},
	Speed:       80,
	W:           26,
	H:           62,
	SpeechColor: theme.InkDim,
})
