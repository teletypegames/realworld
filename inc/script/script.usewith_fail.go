package script

import (
	"math/rand"

	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
	"git.teletypegames.org/games/realworld/inc/world"
)

func UseWithFail(item string, h *inkwell.Hotspot) inkwell.Action {
	key := "fail." + item + "." + h.Name
	world.Game().State.NoteTalked(key)
	n := world.Game().State.Talked(key)

	return inkwell.Seq(
		inkwell.Say(inc.Paul, pick(paulFails, n)),
		inkwell.Say(inc.Dex, dexFail(n)),
	)
}

var paulFails = []inkwell.Text{
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
}

var dexDry = []inkwell.Text{
	{
		{Lang: "en", Text: "No. Not like that."},
		{Lang: "hu", Text: "Nem. Nem így."},
	},
	{
		{Lang: "en", Text: "I heard it. Nothing happened."},
		{Lang: "hu", Text: "Hallottam. Nem történt semmi."},
	},
}

var dexTease = []inkwell.Text{
	{
		{Lang: "en", Text: "Twice the same. The second one rarely goes better."},
		{Lang: "hu", Text: "Kétszer ugyanaz. A második ritkán sikerül jobban."},
	},
	{
		{Lang: "en", Text: "Do it a few more times, maybe physics reconsiders."},
		{Lang: "hu", Text: "Csináld még párszor, hátha a fizika meggondolja magát."},
	},
}

var dexGiveUp = inkwell.Text{
	{Lang: "en", Text: "Paul. Leave it. Look again at what you're carrying — that's where it is."},
	{Lang: "hu", Text: "Paul. Hagyd. Nézd meg még egyszer, mi van nálad — ott lesz."},
}

func dexFail(n int) inkwell.Text {
	switch {
	case n <= 1:
		return pick(dexDry, n)
	case n == 2:
		return pick(dexTease, n)
	default:
		return dexGiveUp
	}
}

func pick(s []inkwell.Text, n int) inkwell.Text {
	switch {
	case len(s) == 0:
		return nil
	case n <= 0:
		return s[0]
	case n-1 < len(s):
		return s[n-1]
	default:
		return s[rand.Intn(len(s))]
	}
}
