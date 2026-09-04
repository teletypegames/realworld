package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var CommentaryTape = ItemManager.Register(&Item{
	Name: "commentary_tape",
	Description: inkwell.Text{
		{Lang: "en", Text: "Noodle's commentary tape"},
		{Lang: "hu", Text: "Noodle kommentárkazettája"},
	},
	OnUseSelf: inkwell.Seq(
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "PAUL, in marker. Underlined twice."},
			{Lang: "hu", Text: "PAUL, filccel. Kétszer aláhúzva."},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "Second slot. He went to some trouble; hear him out."},
			{Lang: "hu", Text: "Kettes foglalat. Fáradt vele; hallgasd végig."},
		}),
	),
})
