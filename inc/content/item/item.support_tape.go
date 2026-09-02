package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = ItemManager.Register(Item{
	Name: inc.ItemSupportTape,
	Description: inkwell.Text{
		{Lang: "en", Text: "“tech support” Personal Tape"},
		{Lang: "hu", Text: "„ügyfélszolgálati” Személyes Kazetta"},
	},
	OnUseSelf: inkwell.Seq(
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Twenty dollars. It was cheap."},
			{Lang: "hu", Text: "Húsz dollár. Olcsó volt."},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "Don't put it in. I'm asking you as a friend. As the only friend you carry around."},
			{Lang: "hu", Text: "Ne tedd be. Barátként kérem. Az egyetlen barátként, akit magaddal hordasz."},
		}),
	),
	OnUseWith: map[string]inkwell.Action{
		inc.ItemMysteryTape: inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "Two tapes, one hand. They don't talk to each other. Nobody does, out of a slot."},
			{Lang: "hu", Text: "Két kazetta, egy kéz. Nem beszélnek egymással. Senki nem beszél, foglalaton kívül."},
		}),
	},
})
