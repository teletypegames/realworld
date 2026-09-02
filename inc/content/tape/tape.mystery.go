package tape

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = TapeManager.Register(Tape{
	Name: inc.TapeMystery,
	Item: inc.ItemMysteryTape,
	Label: inkwell.Text{
		{Lang: "en", Text: "UNKNOWN TAPE"},
		{Lang: "hu", Text: "ISMERETLEN KAZETTA"},
	},
	Color:  theme.Bone,
	Dialog: inc.DialogMysteryTape,
	OnInsert: inkwell.Seq(
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "And now: nothing. Password-locked, top to bottom."},
			{Lang: "hu", Text: "És most: semmi. Jelszóval zárva, tetejétől az aljáig."},
		}),
		inkwell.Say(inc.TapeMystery, inkwell.Text{
			{Lang: "en", Text: "…simulations run on the brain. Guided dreams. What the BVK keeps quiet, it keeps quiet for a reason."},
			{Lang: "hu", Text: "…agyon futtatott szimulációk. Vezérelt álmok. Amit a BVK elhallgat, azt okkal hallgatja el."},
		}),
		inkwell.Say(inc.TapeMystery, inkwell.Text{
			{Lang: "en", Text: "The key forged in the workshop of the underwater sun opens the secrets."},
			{Lang: "hu", Text: "A vízalatti nap műhelyében kovácsoltatott kulcs nyitja a titkokat."},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "That was a riddle. From a tape. I'd like it noted that I never once spoke in riddles."},
			{Lang: "hu", Text: "Ez egy találós kérdés volt. Egy kazettától. Jegyezzük meg, hogy én soha nem beszéltem találós kérdésekben."},
		}),
	),
})
