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

var game *inkwell.Game

func Attach(g *inkwell.Game) {
	game = g
	setMode(ModePlay)
	setNote("")
}

func Game() *inkwell.Game { return game }

func setMode(m Mode) { game.State.SetVar(inc.VarMode, string(m)) }

func setNote(text string) { game.State.SetVar(inc.VarNote, text) }
