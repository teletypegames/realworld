package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Mode string

const (
	ModePlay     Mode = "play"
	ModeCutscene Mode = "cutscene"
	ModeMenu     Mode = "menu"
)

type World struct {
	G *inkwell.Game

	mode Mode

	scene   string
	prev    string
	paused  bool
	pending string
	queue   []inkwell.Action
	running inkwell.Runner
	top     *inkwell.TopBar

	tapeWaiting bool
	tapeLine    inkwell.Action
}

func newWorld(g *inkwell.Game) *World {
	return &World{
		G:    g,
		mode: ModePlay,
	}
}

func (w *World) Mode() Mode { return w.mode }

func (w *World) SetMode(m Mode) { w.mode = m }

func (w *World) HUDVisible() bool { return w.mode == ModePlay }

func (w *World) Scene() string { return w.scene }

func (w *World) Previous() string { return w.prev }

func (w *World) Paused() bool { return w.paused }

func (w *World) SetTopBar(t *inkwell.TopBar) { w.top = t }

func (w *World) Slot2() string {
	if v, ok := w.G.State.Var(VarSlot2).(string); ok {
		return v
	}
	return ""
}

func (w *World) SetSlot2(item string) { w.G.State.SetVar(VarSlot2, item) }

func (w *World) SetPending(item string) { w.pending = item }

func (w *World) TakePending() string {
	item := w.pending
	w.pending = ""
	return item
}

func (w *World) TapeWaiting() bool { return w.tapeWaiting }
