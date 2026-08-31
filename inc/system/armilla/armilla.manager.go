package armilla

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/world"
)

type state struct {
	pending string
	waiting bool
	line    inkwell.Action
}

var st state

func Reset() { st = state{} }

func Slot2() string {
	if v, ok := world.Game().State.Var(inc.VarSlot2).(string); ok {
		return v
	}
	return ""
}

func SetSlot2(item string) { world.Game().State.SetVar(inc.VarSlot2, item) }

func SetPending(item string) { st.pending = item }

func TakePending() string {
	item := st.pending
	st.pending = ""
	return item
}

func SetStrip(text string) { world.Game().State.SetVar(inc.VarArmillaStrip, text) }

func StripIfEmpty(text string) inkwell.Action {
	return world.Fn(func(ctx *inkwell.Ctx) {
		if ctx.Game.State.Var(inc.VarArmillaStrip) == nil {
			ctx.Game.State.SetVar(inc.VarArmillaStrip, text)
		}
	})
}

func Waiting() bool { return st.waiting }

func Offer(a inkwell.Action) inkwell.Action {
	return world.Fn(func(*inkwell.Ctx) {
		st.line = a
		st.waiting = true
	})
}

func TakeOffer() {
	if !st.waiting {
		return
	}
	st.waiting = false
	a := st.line
	st.line = nil
	world.Do(a)
}
