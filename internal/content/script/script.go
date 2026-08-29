// Package script registers the game's named, reusable action sequences —
// cutscenes and recurring logic, built from the inkwell action set.
package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/internal/world"
)

// Register adds every script to the game.
func Register(w *world.World) {
	for _, s := range []inkwell.Script{
		tapeInsert(w),
		awakeningFinale(w),
	} {
		w.G.ScriptManager.Register(s)
	}
}
