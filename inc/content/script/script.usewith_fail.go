package script

import (
	"math/rand"

	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system"
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

func UseWithFail(item string, hotspot *inkwell.Hotspot) inkwell.Action {
	key := "fail." + item + "." + hotspot.Name
	system.System.Game().State.NoteTalked(key)
	count := system.System.Game().State.Talked(key)

	return inkwell.Seq(
		inkwell.Say(inc.Paul, fail.paul(count)),
		inkwell.Say(inc.Dex, fail.dex(count)),
	)
}

func (fail useWithFail) paul(count int) inkwell.Text { return fail.pick(fail.paulLines, count) }

func (fail useWithFail) dex(count int) inkwell.Text {
	switch {
	case count <= 1:
		return fail.pick(fail.dexDry, count)
	case count == 2:
		return fail.pick(fail.dexTease, count)
	default:
		return fail.dexGiveUp
	}
}

func (fail useWithFail) pick(lines []inkwell.Text, count int) inkwell.Text {
	switch {
	case len(lines) == 0:
		return nil
	case count <= 0:
		return lines[0]
	case count-1 < len(lines):
		return lines[count-1]
	default:
		return lines[rand.Intn(len(lines))]
	}
}
