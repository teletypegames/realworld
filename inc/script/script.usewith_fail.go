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

var paulFails = []string{
	"No. That's not going to work.",
	"Tried that. It didn't improve.",
	"All right, I know that one doesn't work.",
	"Now it's personal.",
}

var dexDry = []string{
	"No. Not like that.",
	"I heard it. Nothing happened.",
}

var dexTease = []string{
	"Twice the same. The second one rarely goes better.",
	"Do it a few more times, maybe physics reconsiders.",
}

func dexFail(n int) string {
	switch {
	case n <= 1:
		return pick(dexDry, n)
	case n == 2:
		return pick(dexTease, n)
	default:
		return "Paul. Leave it. Look again at what you're carrying — that's where it is."
	}
}

func pick(s []string, n int) string {
	switch {
	case len(s) == 0:
		return ""
	case n <= 0:
		return s[0]
	case n-1 < len(s):
		return s[n-1]
	default:
		return s[rand.Intn(len(s))]
	}
}
