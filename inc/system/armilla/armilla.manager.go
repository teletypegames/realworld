package armilla

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/world"
)

var (
	pending string

	waiting bool
	line    inkwell.Action
)

func Reset() {
	pending = ""
	waiting = false
	line = nil
}

func Slot2() string {
	if v, ok := world.Game().State.Var(inc.VarSlot2).(string); ok {
		return v
	}
	return ""
}

func SetSlot2(item string) { world.Game().State.SetVar(inc.VarSlot2, item) }

func SetPending(item string) { pending = item }

func TakePending() string {
	item := pending
	pending = ""
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

func Waiting() bool { return waiting }

func Offer(a inkwell.Action) inkwell.Action {
	return world.Fn(func(*inkwell.Ctx) {
		line = a
		waiting = true
	})
}

func TakeOffer() {
	if !waiting {
		return
	}
	waiting = false
	a := line
	line = nil
	world.Do(a)
}
