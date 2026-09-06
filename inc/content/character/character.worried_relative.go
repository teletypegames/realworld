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
	Lore: Text{
		{Lang: "en", Text: "Middle-aged, wringing her coat on the bench. Somebody of hers bought something off the market, and the BVK cares about it far more than a small deal deserves. Would rather not talk about it."},
		{Lang: "hu", Text: "Középkorú, a kabátját gyűrögeti a padon. Valakije vett valamit a piacon, és a BVK sokkal jobban érdeklődik iránta, mint amit egy kis üzlet megérdemel. Nem akar beszélni róla."},
	},
	Speed:       80,
	W:           26,
	H:           62,
	SpeechColor: theme.Bone,
})
