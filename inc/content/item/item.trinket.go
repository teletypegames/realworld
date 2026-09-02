package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = ItemManager.Register(Item{
	Name: inc.ItemTrinket,
	Description: inkwell.Text{
		{Lang: "en", Text: "trinket of unknown purpose"},
		{Lang: "hu", Text: "ismeretlen célú vacak"},
	},
	OnUseSelf: inkwell.Seq(
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Brass. A hinge. A hole where something isn't. Nobody knows what it's for."},
			{Lang: "hu", Text: "Réz. Egy zsanér. Egy lyuk, ahol nincs valami. Senki nem tudja, mire jó."},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "Least of all me, and I've had four seconds longer to think about it than you."},
			{Lang: "hu", Text: "Én a legkevésbé, pedig négy másodperccel többet gondolkodtam rajta, mint te."},
		}),
	),
	OnUseWith: map[string]inkwell.Action{
		inc.ItemSupportTape: inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "The two most useless things you own, together. It's a set now."},
			{Lang: "hu", Text: "A két leghaszontalanabb tulajdonod, együtt. Most már szett."},
		}),
	},
})
