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

	pending string

	tapeWaiting bool
	tapeLine    inkwell.Action
}

var World = &world{}

func (w *world) attach(g *inkwell.Game) {
	w.G = g
	w.pending = ""
	w.tapeWaiting = false
	w.tapeLine = nil
	w.SetMode(ModePlay)
	w.SetNote("")
}

func (w *world) Mode() Mode {
	if v, ok := w.G.State.Var(VarMode).(string); ok {
		return Mode(v)
	}
	return ModePlay
}

func (w *world) SetMode(m Mode) { w.G.State.SetVar(VarMode, string(m)) }

func (w *world) SetNote(text string) { w.G.State.SetVar(VarNote, text) }

func (w *world) Paused() bool { return w.G.State.Var(VarNote) == NotePaused }

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
