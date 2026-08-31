package world

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func Do(action inkwell.Action) { state.game.Do(action) }

type funcAction struct{ body func(*inkwell.Ctx) }

func (action *funcAction) Start() inkwell.Runner {
	return &funcRunner{
		action: action,
	}
}

type funcRunner struct{ action *funcAction }

func (runner *funcRunner) Tick(context *inkwell.Ctx) inkwell.Status {
	runner.action.body(context)
	return inkwell.StatusDone
}

func Func(body func(*inkwell.Ctx)) inkwell.Action {
	return &funcAction{
		body: body,
	}
}

func UseTheme(name string) inkwell.Action {
	return Func(func(context *inkwell.Ctx) { context.Game.UseTheme(name) })
}

func SetMode(mode Mode) inkwell.Action {
	return Func(func(*inkwell.Ctx) { state.setMode(mode) })
}

func Paused(inner inkwell.Action) inkwell.Action {
	return &pausedAction{
		inner: inner,
	}
}

type pausedAction struct {
	inner inkwell.Action
}

func (action *pausedAction) Start() inkwell.Runner {
	return &pausedRunner{
		action: action,
		inner:  action.inner.Start(),
	}
}

type pausedRunner struct {
	action  *pausedAction
	inner   inkwell.Runner
	started bool
}

func (runner *pausedRunner) Tick(context *inkwell.Ctx) inkwell.Status {
	if !runner.started {
		runner.started = true
		state.setNote(inc.NotePaused)
	}
	status := runner.inner.Tick(context)
	if status != inkwell.StatusRunning {
		state.setNote("")
	}
	return status
}
