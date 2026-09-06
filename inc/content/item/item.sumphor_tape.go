package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var SumphorTape = ItemManager.Register(&Item{
	Name: "sumphor_tape",
	Description: Text{
		{Lang: "en", Text: "Personal Tape from Sumphor's desk"},
		{Lang: "hu", Text: "Személyes Kazetta Sumphor asztaláról"},
	},
	Lore: Text{
		{Lang: "en", Text: "The Personal Tape from Sumphor's desk in the abandoned office. Knows the place, not its owner. The unmarked tape says it was Norman's."},
		{Lang: "hu", Text: "A Personal Tape Sumphor asztaláról az elhagyatott irodában. Ismeri a helyet, a gazdáját nem. A jelöletlen kazetta szerint Normané volt."},
	},
	OnUseSelf: inkwell.Seq(
		character.Paul.Say(Text{
			{Lang: "en", Text: "Standard issue, a label somebody peeled off. Dust in the contacts."},
			{Lang: "hu", Text: "Szabványos, a címkét valaki lehúzta. Por az érintkezőkben."},
		}),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "Slot two, if you want to know whose it was. It won't know either, but it'll try."},
			{Lang: "hu", Text: "Kettes foglalat, ha tudni akarod, kié volt. Ő sem fogja tudni, de megpróbálja."},
		}),
	),
})
