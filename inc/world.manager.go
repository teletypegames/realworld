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

type world struct {
	G *inkwell.Game

	mode Mode

	paused  bool
	pending string
	queue   []inkwell.Action
	running inkwell.Runner
	top     *inkwell.TopBar

	tapeWaiting bool
	tapeLine    inkwell.Action
}

var World = &world{}

func (w *world) attach(g *inkwell.Game) {
	w.G = g
	w.mode = ModePlay
	w.paused = false
	w.pending = ""
	w.queue = nil
	w.running = nil
	w.top = nil
	w.tapeWaiting = false
	w.tapeLine = nil
}

func (w *world) Mode() Mode { return w.mode }

func (w *world) SetMode(m Mode) { w.mode = m }

func (w *world) HUDVisible() bool { return w.mode == ModePlay }

func (w *world) Paused() bool { return w.paused }

func (w *world) SetTopBar(t *inkwell.TopBar) { w.top = t }

func (w *world) Slot2() string {
	if v, ok := w.G.State.Var(VarSlot2).(string); ok {
		return v
	}
	return ""
}

func (w *world) SetSlot2(item string) { w.G.State.SetVar(VarSlot2, item) }

func (w *world) SetPending(item string) { w.pending = item }

func (w *world) TakePending() string {
	item := w.pending
	w.pending = ""
	return item
}

func (w *world) TapeWaiting() bool { return w.tapeWaiting }
