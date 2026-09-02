package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = ItemManager.Register(Item{
	Name: inc.ItemSumphorTape,
	Description: inkwell.Text{
		{Lang: "en", Text: "Personal Tape from Sumphor's desk"},
		{Lang: "hu", Text: "Személyes Kazetta Sumphor asztaláról"},
	},
	OnUseSelf: inkwell.Seq(
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Standard issue, a label somebody peeled off. Dust in the contacts."},
			{Lang: "hu", Text: "Szabványos, a címkét valaki lehúzta. Por az érintkezőkben."},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "Slot two, if you want to know whose it was. It won't know either, but it'll try."},
			{Lang: "hu", Text: "Kettes foglalat, ha tudni akarod, kié volt. Ő sem fogja tudni, de megpróbálja."},
		}),
	),
})
