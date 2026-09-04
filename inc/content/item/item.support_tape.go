package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var SupportTape = ItemManager.Register(&Item{
	Name: "support_tape",
	Description: inkwell.Text{
		{Lang: "en", Text: "“tech support” Personal Tape"},
		{Lang: "hu", Text: "„ügyfélszolgálati” Személyes Kazetta"},
	},
	OnUseSelf: inkwell.Seq(
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "Twenty dollars. It was cheap."},
			{Lang: "hu", Text: "Húsz dollár. Olcsó volt."},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "Don't put it in. I'm asking you as a friend. As the only friend you carry around."},
			{Lang: "hu", Text: "Ne tedd be. Barátként kérem. Az egyetlen barátként, akit magaddal hordasz."},
		}),
	),
	OnUseWith: map[*inkwell.Item]inkwell.Action{
		MysteryTape: inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "Two tapes, one hand. They don't talk to each other. Nobody does, out of a slot."},
			{Lang: "hu", Text: "Két kazetta, egy kéz. Nem beszélnek egymással. Senki nem beszél, foglalaton kívül."},
		}),
	},
})
