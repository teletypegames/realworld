package world

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

type Mode string

const (
	ModePlay     Mode = "play"
	ModeCutscene Mode = "cutscene"
	ModeMenu     Mode = "menu"
)

var (
	game *inkwell.Game

	pending string

	tapeWaiting bool
	tapeLine    inkwell.Action
)

func Attach(g *inkwell.Game) {
	game = g
	pending = ""
	tapeWaiting = false
	tapeLine = nil
	setMode(ModePlay)
	setNote("")
}

func Game() *inkwell.Game { return game }

func setMode(m Mode) { game.State.SetVar(inc.VarMode, string(m)) }

func setNote(text string) { game.State.SetVar(inc.VarNote, text) }

func Slot2() string {
	if v, ok := game.State.Var(inc.VarSlot2).(string); ok {
		return v
	}
	return ""
}

func SetSlot2(item string) { game.State.SetVar(inc.VarSlot2, item) }

func SetPending(item string) { pending = item }

func TakePending() string {
	item := pending
	pending = ""
	return item
}

func TapeWaiting() bool { return tapeWaiting }
