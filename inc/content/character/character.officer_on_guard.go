package character

import (
	"realworld/inc/system/theme"
)

var OfficerOnGuard = CharacterManager.Register(&Character{
	Name: "officer_on_guard",
	Label: Text{
		{Lang: "en", Text: "officer"},
		{Lang: "hu", Text: "rendőr"},
	},
	Lore: Text{
		{Lang: "en", Text: "Stands in front of the yellow tape at Noodle's house, arms folded, all day. Nothing to see here, please move along. Hungry, and honest about it."},
		{Lang: "hu", Text: "A sárga szalag előtt áll Noodle házánál, karba tett kézzel, egész nap. Nincs itt semmi látnivaló, tessék továbbmenni. Éhes, és ezt nem is titkolja."},
	},
	Speed:       80,
	W:           34,
	H:           70,
	SpeechColor: theme.InkDim,
})
