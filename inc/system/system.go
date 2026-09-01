package system

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

type Mode string

const (
	ModePlay     Mode = "play"
	ModeCutscene Mode = "cutscene"
)

type system struct {
	game *inkwell.Game
}

var System = &system{}

func (system *system) Attach(game *inkwell.Game) {
	system.game = game
	system.game.State.SetVar(inc.VarMode, string(ModePlay))
	system.game.State.SetVar(inc.VarNote, nil)
}

func (system *system) Game() *inkwell.Game { return system.game }

func (system *system) Do(action inkwell.Action) { system.game.Do(action) }

func (system *system) SetMode(mode Mode) inkwell.Action {
	return inkwell.SetVar(inc.VarMode, string(mode))
}

func (system *system) Paused(inner inkwell.Action) inkwell.Action {
	return inkwell.Noting(inc.VarNote, inc.NotePaused, inner)
}
