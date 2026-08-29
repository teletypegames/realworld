package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Screen struct {
	Name       string
	Title      string
	Background string

	Exits []exit

	OnSelector bool

	Hotspots  []inkwell.Hotspot
	Actors    []inkwell.SceneActor
	Walkboxes []inkwell.Polygon
	OnEnter   inkwell.Action
}

var ScreenManager = NewManager(func(entity Screen) string { return entity.Name })

func registerScreen() {
	fillSelectorPins()
	registerAll(ScreenManager, func(entity Screen) {
		World.G.SceneManager.Register(screenBuild(entity))
	})
}

var screenFloor = []inkwell.Polygon{
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

func screenBuild(s Screen) inkwell.Scene {
	exits := s.Exits
	if s.OnSelector {
		exits = append(exits, toSelector())
	}
	hotspots := make([]inkwell.Hotspot, 0, len(exits)+len(s.Hotspots))
	hotspots = append(hotspots, s.Hotspots...)
	for _, e := range exits {
		hotspots = append(hotspots, e.hotspot())
	}

	actors := s.Actors
	if actors == nil {
		actors = []inkwell.SceneActor{
			{
				CharacterName: Paul,
				At: inkwell.Point{
					X: 320,
					Y: 232,
				},
			},
		}
	}
	walkboxes := s.Walkboxes
	if walkboxes == nil {
		walkboxes = screenFloor
	}
	onEnter := inkwell.Action(EnterScene(s.Name))
	if s.OnEnter != nil {
		onEnter = inkwell.Seq(onEnter, s.OnEnter)
	}

	return inkwell.Scene{
		Name:       s.Name,
		Title:      s.Title,
		Background: s.Background,
		Actors:     actors,
		Walkboxes:  walkboxes,
		Hotspots:   hotspots,
		OnEnter:    onEnter,
	}
}

func setVarIfEmpty(name string, v any) inkwell.Action {
	return Fn(func(ctx *inkwell.Ctx) {
		if ctx.Game.State.Var(name) == nil {
			ctx.Game.State.SetVar(name, v)
		}
	})
}
