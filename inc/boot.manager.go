// Package inc is the whole game, in one flat package.
//
// Every file is named [category].[name].go: one category per concern, one
// entity per file. A category's <category>.manager.go holds what ties that
// category together — its Register, its shared types — and the files beside it
// hold one screen, one background, one item each.
//
// This file is the boot category: it wires the layers together, and it is the
// only place where the world, the content, the theme and the HUD meet. main.go
// knows nothing else.
package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// Opts are the run parameters.
type Opts struct {
	Screen string // starting screen; empty means the alley
	Finale bool   // start with the finale, to try the Nokia-punk theme switch
}

// New builds a ready-to-run game.
func New(o Opts) *inkwell.Game {
	start := o.Screen
	if start == "" {
		start = ScreenAlley
	}

	g := inkwell.NewGame("Real World", ScreenW, ScreenH)
	g.MaxLogLines = 64

	w := newWorld(g)
	registerTheme(g)
	registerContent(w)

	g.UseTheme(RealWorld)
	registerUI(w)

	g.StartAt(start)
	g.OnStart(inkwell.Seq(bootOpening(w, start, o.Finale)...))
	return g
}

func bootOpening(w *World, start string, finale bool) []inkwell.Action {
	acts := []inkwell.Action{
		EnterScene(w, start),
		// Beat 2 (the police station) will set the tip flag. Until that beat
		// is written we set it here so the alley slice is playable end to end —
		// the conditional branch itself is real code, not bypassed.
		inkwell.SetFlag(FlagPoliceTip),
		inkwell.Say(Paul, "Back alley. Just like the clerk said."),
		TapeSay(Dex, "A government office tipped you off to remove something from a scene. That doesn't strike you as odd?"),
	}
	if finale {
		acts = append(acts, inkwell.RunScript(ScriptFinale))
	}
	return acts
}
