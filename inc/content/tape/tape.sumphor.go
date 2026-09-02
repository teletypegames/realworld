package tape

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = TapeManager.Register(Tape{
	Name: inc.TapeSumphor,
	Item: inc.ItemSumphorTape,
	Label: inkwell.Text{
		{Lang: "en", Text: "SUMPHOR'S TAPE"},
		{Lang: "hu", Text: "SUMPHOR KAZETTÁJA"},
	},
	Color:  theme.InkDim,
	Dialog: inc.DialogSumphorTape,
	OnInsert: inkwell.Seq(
		inkwell.Say(inc.TapeSumphor, inkwell.Text{
			{Lang: "en", Text: "Desk four. Second floor. Coffee at ten. Who is this?"},
			{Lang: "hu", Text: "Négyes asztal. Második emelet. Kávé tízkor. Ki vagy te?"},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "It knows the building better than it knows its own owner. That's a kind of damage."},
			{Lang: "hu", Text: "Az épületet jobban ismeri, mint a saját gazdáját. Ez is egyfajta sérülés."},
		}),
	),
})
