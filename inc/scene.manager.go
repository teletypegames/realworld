package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Scene = inkwell.Scene

var SceneManager = inkwell.NewManager[Scene]()

func registerScene() {
	fillSelectorPins()
	SceneManager.Each(func(entity Scene) {
		World.G.SceneManager.Register(sceneDefaults(entity))
	})
}

var sceneFloor = []inkwell.Polygon{
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

func sceneDefaults(s Scene) Scene {
	if s.Actors == nil {
		s.Actors = []inkwell.SceneActor{
			{
				CharacterName: Paul,
				At: inkwell.Point{
					X: 320,
					Y: 232,
				},
			},
		}
	}
	if s.Walkboxes == nil {
		s.Walkboxes = sceneFloor
	}
	return s
}

func setVarIfEmpty(name string, v any) inkwell.Action {
	return Fn(func(ctx *inkwell.Ctx) {
		if ctx.Game.State.Var(name) == nil {
			ctx.Game.State.SetVar(name, v)
		}
	})
}
