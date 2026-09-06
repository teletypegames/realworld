package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var IceCream = ItemManager.Register(&Item{
	Name: "ice_cream",
	Description: Text{
		{Lang: "en", Text: "pistachio ice cream"},
		{Lang: "hu", Text: "pisztáciafagyi"},
	},
	Lore: Text{
		{Lang: "en", Text: "Pistachio, melting. Everything in this town is either melting or pretending it never existed."},
		{Lang: "hu", Text: "Pisztácia, olvadóban. Ebben a városban minden vagy olvad, vagy úgy tesz, mintha sosem létezett volna."},
	},
	OnUseSelf: inkwell.Seq(
		character.Paul.Say(Text{
			{Lang: "en", Text: "It's melting. Everything in this town is either melting or pretending it never existed."},
			{Lang: "hu", Text: "Olvad. Ebben a városban minden vagy olvad, vagy úgy tesz, mintha soha nem létezett volna."},
		}),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "Eat it or give it away. Don't philosophise at a cone."},
			{Lang: "hu", Text: "Edd meg vagy add oda. Ne filozofálj egy tölcsérnek."},
		}),
	),
	OnUseWith: map[*inkwell.Item]Action{
		BlackArmilla: speaker.Dex.Say(Text{
			{Lang: "en", Text: "Cold and sticky on a bracelet that's already broken. Bold."},
			{Lang: "hu", Text: "Hideg és ragacsos egy karperecen, ami már így is rossz. Merész."},
		}),
	},
})
