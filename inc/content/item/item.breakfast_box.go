package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var BreakfastBox = ItemManager.Register(&Item{
	Name: "breakfast_box",
	Description: inkwell.Text{
		{Lang: "en", Text: "box of breakfast"},
		{Lang: "hu", Text: "doboz reggeli"},
	},
	OnUseSelf: inkwell.Seq(
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "One dumpling won't hurt."},
			{Lang: "hu", Text: "Egy gombóc nem árt meg."},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "Save it. Somebody out there is hungrier than you, and more useful."},
			{Lang: "hu", Text: "Tedd el. Valaki odakint éhesebb nálad, és hasznosabb is."},
		}),
	),
})
