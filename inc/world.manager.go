package inc

// World — the domain aggregate: a thin layer around the inkwell Game
// holding the runtime state that must NOT be saved, plus the game's own
// actions and the action pump.
//
// The world knows nothing about the HUD or the content: the ui and content
// categories both build on it, never the other way round.

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// Mode is the HUD mode.
//
// IMPORTANT: this is deliberately NOT a State.Var. inkwell's state.save.go
// serialises State.Vars but drops the in-flight script ("saves capture state at
// idle/ready boundaries"), so a save taken mid-cutscene would reload into a
// permanently letterboxed, HUD-less game with no script running. The mode is
// therefore an unsaved runtime field.
type Mode string

const (
	ModePlay     Mode = "play"     // full HUD
	ModeCutscene Mode = "cutscene" // letterbox, HUD hidden
	ModeMenu     Mode = "menu"     // main menu, HUD hidden
)

// World holds the runtime state that does not belong in a save file.
type World struct {
	G *inkwell.Game

	mode Mode

	scene   string           // current scene — inkwell exposes no accessor for it
	prev    string           // the one before it, for screens with no fixed way out
	paused  bool             // is the world stopped (during dialogue), for the TopBar suffix
	pending string           // the tape currently being loaded into slot 2
	queue   []inkwell.Action // HUD-originated actions waiting to run
	running inkwell.Runner
	top     *inkwell.TopBar // the location bar, so the pump can write its suffix

	// During a cutscene a tape signals but never interrupts: the line waits
	// here until the player asks for it with SPACE.
	tapeWaiting bool
	tapeLine    inkwell.Action
}

// New wraps an existing inkwell game.
func newWorld(g *inkwell.Game) *World {
	return &World{G: g, mode: ModePlay}
}

// Mode returns the current HUD mode.
func (w *World) Mode() Mode { return w.mode }

// SetMode switches the HUD mode.
func (w *World) SetMode(m Mode) { w.mode = m }

// HUDVisible reports whether the HUD should draw: only in play mode.
func (w *World) HUDVisible() bool { return w.mode == ModePlay }

// Scene returns the name of the current scene.
func (w *World) Scene() string { return w.scene }

// Previous returns the screen we came from, empty at the start of the game.
func (w *World) Previous() string { return w.prev }

// Paused reports whether the world is stopped (during dialogue).
func (w *World) Paused() bool { return w.paused }

// SetTopBar hands over the location bar so the pump can write its suffix.
func (w *World) SetTopBar(t *inkwell.TopBar) { w.top = t }

// Slot2 returns the item name in the second tape slot, empty if there is none.
//
// Unlike Mode, this is REAL game state and lives in State.Var so that it is
// saved.
func (w *World) Slot2() string {
	if v, ok := w.G.State.Var(VarSlot2).(string); ok {
		return v
	}
	return ""
}

// SetSlot2 sets the contents of the second tape slot.
func (w *World) SetSlot2(item string) { w.G.State.SetVar(VarSlot2, item) }

// SetPending marks which tape is going into slot 2. The tape-insert script
// performs the actual load; the widget only marks.
func (w *World) SetPending(item string) { w.pending = item }

// TakePending returns and clears the tape waiting to be loaded.
func (w *World) TakePending() string {
	item := w.pending
	w.pending = ""
	return item
}

// TapeWaiting reports whether a tape line has been offered but not yet spoken.
func (w *World) TapeWaiting() bool { return w.tapeWaiting }
