package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var Trinket = ItemManager.Register(&Item{
	Name: "trinket",
	Description: Text{
		{Lang: "en", Text: "trinket of unknown purpose"},
		{Lang: "hu", Text: "ismeretlen célú vacak"},
	},
	OnUseSelf: inkwell.Seq(
		character.Paul.Say(Text{
			{Lang: "en", Text: "Brass. A hinge. A hole where something isn't. Nobody knows what it's for."},
			{Lang: "hu", Text: "Réz. Egy zsanér. Egy lyuk, ahol nincs valami. Senki nem tudja, mire jó."},
		}),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "Least of all me, and I've had four seconds longer to think about it than you."},
			{Lang: "hu", Text: "Én a legkevésbé, pedig négy másodperccel többet gondolkodtam rajta, mint te."},
		}),
	),
	OnUseWith: map[*inkwell.Item]Action{
		SupportTape: speaker.Dex.Say(Text{
			{Lang: "en", Text: "The two most useless things you own, together. It's a set now."},
			{Lang: "hu", Text: "A két leghaszontalanabb tulajdonod, együtt. Most már szett."},
		}),
	},
})
