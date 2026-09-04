package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var MysteryTape = ItemManager.Register(&Item{
	Name: "mystery_tape",
	Description: inkwell.Text{
		{Lang: "en", Text: "unmarked Personal Tape"},
		{Lang: "hu", Text: "jelöletlen Személyes Kazetta"},
	},
	OnUseSelf: inkwell.Seq(
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "Password-locked. Of course."},
			{Lang: "hu", Text: "Jelszóval zárva. Persze."},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "Put it in the second slot if you care that much. I did warn you."},
			{Lang: "hu", Text: "Tedd a második foglalatba, ha ennyire érdekel. Én szóltam."},
		}),
	),
})
