package system

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

type Mode string

const (
	ModePlay     Mode = "play"
	ModeCutscene Mode = "cutscene"
	ModeMenu     Mode = "menu"
)

type system struct {
	game *inkwell.Game
}

var System = &system{}

func (system *system) Attach(game *inkwell.Game) {
	system.game = game
	system.setMode(ModePlay)
	system.setNote("")
}

func (system *system) Game() *inkwell.Game { return system.game }

func (system *system) Do(action inkwell.Action) { system.game.Do(action) }

func (system *system) Func(body func(*inkwell.Ctx)) inkwell.Action {
	return inkwell.Custom(func(context *inkwell.Ctx) inkwell.Status {
		body(context)
		return inkwell.StatusDone
	})
}

func (system *system) SetMode(mode Mode) inkwell.Action {
	return system.Func(func(*inkwell.Ctx) { system.setMode(mode) })
}

func (system *system) UseTheme(name string) inkwell.Action {
	return system.Func(func(context *inkwell.Ctx) { context.Game.UseTheme(name) })
}

func (system *system) Paused(inner inkwell.Action) inkwell.Action {
	var runner inkwell.Runner
	return inkwell.Custom(func(context *inkwell.Ctx) inkwell.Status {
		if runner == nil {
			runner = inner.Start()
			system.setNote(inc.NotePaused)
		}
		status := runner.Tick(context)
		if status != inkwell.StatusRunning {
			runner = nil
			system.setNote("")
		}
		return status
	})
}

func (system *system) setMode(mode Mode) { system.game.State.SetVar(inc.VarMode, string(mode)) }

func (system *system) setNote(note string) { system.game.State.SetVar(inc.VarNote, note) }
