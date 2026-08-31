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

type state struct {
	game *inkwell.Game
}

var st state

func Attach(g *inkwell.Game) {
	st = state{
		game: g,
	}
	st.setMode(ModePlay)
	st.setNote("")
}

func Game() *inkwell.Game { return st.game }

func (s *state) setMode(m Mode) { s.game.State.SetVar(inc.VarMode, string(m)) }

func (s *state) setNote(text string) { s.game.State.SetVar(inc.VarNote, text) }
