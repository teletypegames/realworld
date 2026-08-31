package armilla

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/world"
)

type armillaState struct {
	pending string
	waiting bool
	line    inkwell.Action
}

var state armillaState

func Reset() { state = armillaState{} }

func Slot2() string {
	if value, isText := world.Game().State.Var(inc.VarSlot2).(string); isText {
		return value
	}
	return ""
}

func SetSlot2(item string) { world.Game().State.SetVar(inc.VarSlot2, item) }

func SetPending(item string) { state.pending = item }

func TakePending() string {
	item := state.pending
	state.pending = ""
	return item
}

func SetStrip(text string) { world.Game().State.SetVar(inc.VarArmillaStrip, text) }

func StripIfEmpty(text string) inkwell.Action {
	return world.Func(func(context *inkwell.Ctx) {
		if context.Game.State.Var(inc.VarArmillaStrip) == nil {
			context.Game.State.SetVar(inc.VarArmillaStrip, text)
		}
	})
}

func Waiting() bool { return state.waiting }

func Offer(line inkwell.Action) inkwell.Action {
	return world.Func(func(*inkwell.Ctx) {
		state.line = line
		state.waiting = true
	})
}

func TakeOffer() {
	if !state.waiting {
		return
	}
	state.waiting = false
	line := state.line
	state.line = nil
	world.Do(line)
}
