package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var CommentaryTape = ItemManager.Register(&Item{
	Name: "commentary_tape",
	Description: Text{
		{Lang: "en", Text: "Noodle's commentary tape"},
		{Lang: "hu", Text: "Noodle kommentárkazettája"},
	},
	Lore: Text{
		{Lang: "en", Text: "PAUL, in marker, underlined twice. Noodle's developer commentary, the tape the police carried out of his back room before Paul got there."},
		{Lang: "hu", Text: "PAUL, filccel, kétszer aláhúzva. Noodle fejlesztői kommentárja, a kazetta, amit a rendőrség kivitt a hátsó szobájából, mielőtt Paul odaért."},
	},
	OnUseSelf: inkwell.Seq(
		character.Paul.Say(Text{
			{Lang: "en", Text: "PAUL, in marker. Underlined twice."},
			{Lang: "hu", Text: "PAUL, filccel. Kétszer aláhúzva."},
		}),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "Second slot. He went to some trouble; hear him out."},
			{Lang: "hu", Text: "Kettes foglalat. Fáradt vele; hallgasd végig."},
		}),
	),
})
