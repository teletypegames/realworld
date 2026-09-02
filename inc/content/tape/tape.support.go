package tape

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = TapeManager.Register(Tape{
	Name: inc.TapeSupport,
	Item: inc.ItemSupportTape,
	Label: inkwell.Text{
		{Lang: "en", Text: "TECH SUPPORT"},
		{Lang: "hu", Text: "ÜGYFÉLSZOLGÁLAT"},
	},
	Color:  theme.AmberPale,
	Dialog: inc.DialogSupportTape,
	OnInsert: inkwell.Seq(
		inkwell.Say(inc.TapeSupport, inkwell.Text{
			{Lang: "en", Text: "Thank you for calling. My name is Raj. Have you tried turning it off and on again?"},
			{Lang: "hu", Text: "Köszönjük a hívását. A nevem Raj. Próbálta már ki- és bekapcsolni?"},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "Paul. Take it out. Take it out now."},
			{Lang: "hu", Text: "Paul. Vedd ki. Most vedd ki."},
		}),
	),
})
