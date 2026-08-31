package world

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

type worldState struct {
	game *inkwell.Game
}

var state worldState

func Attach(game *inkwell.Game) {
	state = worldState{
		game: game,
	}
	state.setMode(ModePlay)
	state.setNote("")
}

func Game() *inkwell.Game { return state.game }

func (state *worldState) setMode(mode Mode) { state.game.State.SetVar(inc.VarMode, string(mode)) }

func (state *worldState) setNote(note string) { state.game.State.SetVar(inc.VarNote, note) }
