package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = ItemManager.Register(Item{
	Name: inc.ItemCommentaryTape,
	Description: inkwell.Text{
		{Lang: "en", Text: "Noodle's commentary tape"},
		{Lang: "hu", Text: "Noodle kommentárkazettája"},
	},
	OnUseSelf: inkwell.Seq(
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "PAUL, in marker. Underlined twice."},
			{Lang: "hu", Text: "PAUL, filccel. Kétszer aláhúzva."},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "Second slot. He went to some trouble; hear him out."},
			{Lang: "hu", Text: "Kettes foglalat. Fáradt vele; hallgasd végig."},
		}),
	),
})
