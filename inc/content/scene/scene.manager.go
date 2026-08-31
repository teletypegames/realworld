package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/world"
)

type Scene = inkwell.Scene

var SceneManager = inkwell.NewManager[Scene]()

var Floor = []inkwell.Polygon{
	inkwell.Poly(
		inkwell.Point{
			X: 16,
			Y: 196,
		}, inkwell.Point{
			X: 624,
			Y: 196,
		},
		inkwell.Point{
			X: 624,
			Y: 258,
		}, inkwell.Point{
			X: 16,
			Y: 258,
		},
	),
}

func setVarIfEmpty(name string, v any) inkwell.Action {
	return world.Fn(func(ctx *inkwell.Ctx) {
		if ctx.Game.State.Var(name) == nil {
			ctx.Game.State.SetVar(name, v)
		}
	})
}
