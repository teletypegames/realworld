package tape

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/dialog"
	"realworld/inc/content/item"
	"realworld/inc/content/speaker"
)

var Support = TapeManager.Register(&Tape{
	Name:    "tape_support",
	Item:    item.SupportTape,
	Speaker: speaker.TapeSupport,
	Dialog:  dialog.SupportTape,
	OnInsert: inkwell.Seq(
		speaker.TapeSupport.Say(Text{
			{Lang: "en", Text: "Thank you for calling. My name is Raj. Have you tried turning it off and on again?"},
			{Lang: "hu", Text: "Köszönjük a hívását. A nevem Raj. Próbálta már ki- és bekapcsolni?"},
		}),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "Paul. Take it out. Take it out now."},
			{Lang: "hu", Text: "Paul. Vedd ki. Most vedd ki."},
		}),
	),
})
