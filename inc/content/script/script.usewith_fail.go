package script

import (
	"math/rand"

	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/world"
)

type useWithFail struct {
	paulLines []inkwell.Text
	dexDry    []inkwell.Text
	dexTease  []inkwell.Text
	dexGiveUp inkwell.Text
}

var fail = useWithFail{
	paulLines: []inkwell.Text{
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
	dexDry: []inkwell.Text{
		{
			{Lang: "en", Text: "No. Not like that."},
			{Lang: "hu", Text: "Nem. Nem így."},
		},
		{
			{Lang: "en", Text: "I heard it. Nothing happened."},
			{Lang: "hu", Text: "Hallottam. Nem történt semmi."},
		},
	},
	dexTease: []inkwell.Text{
		{
			{Lang: "en", Text: "Twice the same. The second one rarely goes better."},
			{Lang: "hu", Text: "Kétszer ugyanaz. A második ritkán sikerül jobban."},
		},
		{
			{Lang: "en", Text: "Do it a few more times, maybe physics reconsiders."},
			{Lang: "hu", Text: "Csináld még párszor, hátha a fizika meggondolja magát."},
		},
	},
	dexGiveUp: inkwell.Text{
		{Lang: "en", Text: "Paul. Leave it. Look again at what you're carrying — that's where it is."},
		{Lang: "hu", Text: "Paul. Hagyd. Nézd meg még egyszer, mi van nálad — ott lesz."},
	},
}

func UseWithFail(item string, h *inkwell.Hotspot) inkwell.Action {
	key := "fail." + item + "." + h.Name
	world.Game().State.NoteTalked(key)
	n := world.Game().State.Talked(key)

	return inkwell.Seq(
		inkwell.Say(inc.Paul, fail.paul(n)),
		inkwell.Say(inc.Dex, fail.dex(n)),
	)
}

func (f useWithFail) paul(n int) inkwell.Text { return f.pick(f.paulLines, n) }

func (f useWithFail) dex(n int) inkwell.Text {
	switch {
	case n <= 1:
		return f.pick(f.dexDry, n)
	case n == 2:
		return f.pick(f.dexTease, n)
	default:
		return f.dexGiveUp
	}
}

func (f useWithFail) pick(s []inkwell.Text, n int) inkwell.Text {
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
