package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var MysteryTape = ItemManager.Register(&Item{
	Name: "mystery_tape",
	Description: Text{
		{Lang: "en", Text: "unmarked Personal Tape"},
		{Lang: "hu", Text: "jelöletlen Személyes Kazetta"},
	},
	Lore: Text{
		{Lang: "en", Text: "An unmarked Personal Tape from behind Noodle's house. Password-locked, until it wasn't. Talks about a project it isn't supposed to know, and about Norman, whom it says it never met."},
		{Lang: "hu", Text: "Jelöletlen Personal Tape Noodle háza mögül. Jelszóval zárt, amíg nem volt az. Egy projektről beszél, amiről nem kéne tudnia, és Normanról, akivel állítása szerint sosem találkozott."},
	},
	OnUseSelf: inkwell.Seq(
		character.Paul.Say(Text{
			{Lang: "en", Text: "Password-locked. Of course."},
			{Lang: "hu", Text: "Jelszóval zárva. Persze."},
		}),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "Put it in the second slot if you care that much. I did warn you."},
			{Lang: "hu", Text: "Tedd a második foglalatba, ha ennyire érdekel. Én szóltam."},
		}),
	),
})
