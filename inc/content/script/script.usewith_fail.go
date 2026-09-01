package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system"
)

type useWithFail struct {
	paul inkwell.Ladder
	dex  inkwell.Ladder
}

var fail = useWithFail{
	paul: inkwell.Ladder{
		Rungs: []inkwell.Text{
			{
				{Lang: "en", Text: "No. That's not going to work."},
				{Lang: "hu", Text: "Nem. Ez így nem fog menni."},
			},
			{
				{Lang: "en", Text: "Tried that. It didn't improve."},
				{Lang: "hu", Text: "Ezt már próbáltam. Nem lett jobb."},
			},
			{
				{Lang: "en", Text: "All right, I know that one doesn't work."},
				{Lang: "hu", Text: "Jó, tudom, hogy az nem működik."},
			},
			{
				{Lang: "en", Text: "Now it's personal."},
				{Lang: "hu", Text: "Most már személyes."},
			},
		},
	},
	dex: inkwell.Ladder{
		Rungs: []inkwell.Text{
			{
				{Lang: "en", Text: "No. Not like that."},
				{Lang: "hu", Text: "Nem. Nem így."},
			},
			{
				{Lang: "en", Text: "I heard it. Nothing happened."},
				{Lang: "hu", Text: "Hallottam. Nem történt semmi."},
			},
			{
				{Lang: "en", Text: "Twice the same. The second one rarely goes better."},
				{Lang: "hu", Text: "Kétszer ugyanaz. A második ritkán sikerül jobban."},
			},
			{
				{Lang: "en", Text: "Do it a few more times, maybe physics reconsiders."},
				{Lang: "hu", Text: "Csináld még párszor, hátha a fizika meggondolja magát."},
			},
		},
		Last: inkwell.Text{
			{Lang: "en", Text: "Paul. Leave it. Look again at what you're carrying — that's where it is."},
			{Lang: "hu", Text: "Paul. Hagyd. Nézd meg még egyszer, mi van nálad — ott lesz."},
		},
	},
}

func UseWithFail(item string, hotspot *inkwell.Hotspot) inkwell.Action {
	state := system.System.Game().State
	key := "fail." + item + "." + hotspot.Name
	state.NoteTalked(key)
	count := state.Talked(key)

	return inkwell.Seq(
		inkwell.Say(inc.Paul, fail.paul.At(count)),
		inkwell.Say(inc.Dex, fail.dex.At(count)),
	)
}
