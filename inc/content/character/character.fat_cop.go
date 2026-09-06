package character

import (
	"realworld/inc/system/theme"
)

var FatCop = CharacterManager.Register(&Character{
	Name: "fat_cop",
	Label: Text{
		{Lang: "en", Text: "fat cop"},
		{Lang: "hu", Text: "kövér rendőr"},
	},
	Lore: Text{
		{Lang: "en", Text: "Leans on the cruiser outside the station and works through a donut. Tells the thin figure in the passenger seat everything he shouldn't: Sumphor, Noodle, the odd types in Chinatown."},
		{Lang: "hu", Text: "A kapitányság előtt a járőrkocsinak dőlve eszi a fánkját. Mindent elmond az anyósülésen ülő vékony alaknak, amit nem kéne: Sumphort, Noodle-t, a kínai negyed fura figuráit."},
	},
	Speed:       80,
	W:           40,
	H:           64,
	SpeechColor: theme.InkDim,
})
