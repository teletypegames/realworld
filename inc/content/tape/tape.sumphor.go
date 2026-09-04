package tape

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/dialog"
	"realworld/inc/content/item"
	"realworld/inc/content/speaker"
)

var Sumphor = TapeManager.Register(&Tape{
	Name:    "tape_sumphor",
	Item:    item.SumphorTape,
	Speaker: speaker.TapeSumphor,
	Dialog:  dialog.SumphorTape,
	OnInsert: inkwell.Seq(
		inkwell.Say(speaker.TapeSumphor, inkwell.Text{
			{Lang: "en", Text: "Desk four. Second floor. Coffee at ten. Who is this?"},
			{Lang: "hu", Text: "Négyes asztal. Második emelet. Kávé tízkor. Ki vagy te?"},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "It knows the building better than it knows its own owner. That's a kind of damage."},
			{Lang: "hu", Text: "Az épületet jobban ismeri, mint a saját gazdáját. Ez is egyfajta sérülés."},
		}),
	),
})
