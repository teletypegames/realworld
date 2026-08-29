package inc

// Scripts — the game's named, reusable action sequences:
// cutscenes and recurring logic, built from the inkwell action set.

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// registerScript adds every script to the game.
func registerScript(w *World) {
	for _, s := range []inkwell.Script{
		scriptTapeInsert(w),
		scriptAwakeningFinale(w),
	} {
		w.G.ScriptManager.Register(s)
	}
}
