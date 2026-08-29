// Package scene registers the game's scenes.
package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/world"
)

// Register adds every scene to the game.
func Register(w *world.World) {
	for _, s := range []inkwell.Scene{
		alley(w),
	} {
		w.G.SceneManager.Register(s)
	}
}

// setVarIfEmpty only writes when the variable is still unset, so re-entering a
// scene does not clobber what the game has set in the meantime.
func setVarIfEmpty(name string, v any) inkwell.Action {
	return world.Fn(func(ctx *inkwell.Ctx) {
		if ctx.Game.State.Var(name) == nil {
			ctx.Game.State.SetVar(name, v)
		}
	})
}
