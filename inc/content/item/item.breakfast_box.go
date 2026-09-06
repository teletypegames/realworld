package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var BreakfastBox = ItemManager.Register(&Item{
	Name: "breakfast_box",
	Description: Text{
		{Lang: "en", Text: "box of breakfast"},
		{Lang: "hu", Text: "doboz reggeli"},
	},
	Lore: Text{
		{Lang: "en", Text: "Dumplings in a box from the airport. A hungry officer's price for an address."},
		{Lang: "hu", Text: "Gombócok egy dobozban a reptérről. Egy éhes rendőr ára egy címért."},
	},
	OnUseSelf: inkwell.Seq(
		character.Paul.Say(Text{
			{Lang: "en", Text: "One dumpling won't hurt."},
			{Lang: "hu", Text: "Egy gombóc nem árt meg."},
		}),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "Save it. Somebody out there is hungrier than you, and more useful."},
			{Lang: "hu", Text: "Tedd el. Valaki odakint éhesebb nálad, és hasznosabb is."},
		}),
	),
})
