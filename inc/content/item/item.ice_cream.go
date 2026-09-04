package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var IceCream = ItemManager.Register(&Item{
	Name: "ice_cream",
	Description: inkwell.Text{
		{Lang: "en", Text: "pistachio ice cream"},
		{Lang: "hu", Text: "pisztáciafagyi"},
	},
	OnUseSelf: inkwell.Seq(
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "It's melting. Everything in this town is either melting or pretending it never existed."},
			{Lang: "hu", Text: "Olvad. Ebben a városban minden vagy olvad, vagy úgy tesz, mintha soha nem létezett volna."},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "Eat it or give it away. Don't philosophise at a cone."},
			{Lang: "hu", Text: "Edd meg vagy add oda. Ne filozofálj egy tölcsérnek."},
		}),
	),
	OnUseWith: map[*inkwell.Item]inkwell.Action{
		BlackArmilla: inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "Cold and sticky on a bracelet that's already broken. Bold."},
			{Lang: "hu", Text: "Hideg és ragacsos egy karperecen, ami már így is rossz. Merész."},
		}),
	},
})
