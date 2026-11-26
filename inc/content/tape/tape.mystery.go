package tape

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/dialog"
	"realworld/inc/content/item"
	"realworld/inc/content/speaker"
)

var Mystery = TapeManager.Register(&Tape{
	Name:    "tape_mystery",
	Item:    item.MysteryTape,
	Speaker: speaker.TapeMystery,
	Dialog:  dialog.MysteryTape,
	OnInsert: inkwell.Seq(
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "And now: nothing. Password-locked, top to bottom."},
			{Lang: "hu", Text: "És most: semmi. Jelszóval zárva, tetejétől az aljáig."},
		}),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "System alert: Digital wallet locked. Trigger: unmarked tape. Access to digital funds has been revoked."},
			{Lang: "hu", Text: "Rendszerüzenet: Digitális pénztárca zárolva. Kiváltó ok: ismeretlen kazetta. A digitális pénzeszközökhöz való hozzáférés megszűnt."},
		}),
		speaker.TapeMystery.Say(Text{
			{Lang: "en", Text: "…simulations run on the brain. Guided dreams. What the BVK keeps quiet, it keeps quiet for a reason."},
			{Lang: "hu", Text: "…agyon futtatott szimulációk. Vezérelt álmok. Amit a BVK elhallgat, azt okkal hallgatja el."},
		}),
		speaker.TapeMystery.Say(Text{
			{Lang: "en", Text: "The key forged in the workshop of the underwater sun opens the secrets."},
			{Lang: "hu", Text: "A vízalatti nap műhelyében kovácsoltatott kulcs nyitja a titkokat."},
		}),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "That was a riddle. From a tape. I'd like it noted that I never once spoke in riddles."},
			{Lang: "hu", Text: "Ez egy találós kérdés volt. Egy kazettától. Jegyezzük meg, hogy én soha nem beszéltem találós kérdésekben."},
		}),
	),
})
