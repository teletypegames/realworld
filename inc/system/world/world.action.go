package world

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func Do(a inkwell.Action) { game.Do(a) }

type fnAction struct{ fn func(*inkwell.Ctx) }

func (a *fnAction) Start() inkwell.Runner {
	return &fnRunner{
		spec: a,
	}
}

type fnRunner struct{ spec *fnAction }

func (r *fnRunner) Tick(ctx *inkwell.Ctx) inkwell.Status {
	r.spec.fn(ctx)
	return inkwell.StatusDone
}

func Fn(f func(*inkwell.Ctx)) inkwell.Action {
	return &fnAction{
		fn: f,
	}
}

func UseTheme(name string) inkwell.Action {
	return Fn(func(ctx *inkwell.Ctx) { ctx.Game.UseTheme(name) })
}

func SetMode(m Mode) inkwell.Action {
	return Fn(func(*inkwell.Ctx) { setMode(m) })
}

func Paused(inner inkwell.Action) inkwell.Action {
	return &pausedAction{
		inner: inner,
	}
}

type pausedAction struct {
	inner inkwell.Action
}

func (a *pausedAction) Start() inkwell.Runner {
	return &pausedRunner{
		spec:  a,
		inner: a.inner.Start(),
	}
}

type pausedRunner struct {
	spec    *pausedAction
	inner   inkwell.Runner
	started bool
}

func (r *pausedRunner) Tick(ctx *inkwell.Ctx) inkwell.Status {
	if !r.started {
		r.started = true
		setNote(inc.NotePaused)
	}
	s := r.inner.Tick(ctx)
	if s != inkwell.StatusRunning {
		setNote("")
	}
	return s
}
