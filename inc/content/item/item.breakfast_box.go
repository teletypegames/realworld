package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = ItemManager.Register(Item{
	Name: inc.ItemBreakfastBox,
	Description: inkwell.Text{
		{Lang: "en", Text: "box of breakfast"},
		{Lang: "hu", Text: "doboz reggeli"},
	},
	OnUseSelf: inkwell.Seq(
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "One dumpling won't hurt."},
			{Lang: "hu", Text: "Egy gombóc nem árt meg."},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "Save it. Somebody out there is hungrier than you, and more useful."},
			{Lang: "hu", Text: "Tedd el. Valaki odakint éhesebb nálad, és hasznosabb is."},
		}),
	),
})
