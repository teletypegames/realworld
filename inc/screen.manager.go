package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type screen struct {
	name       string
	title      string
	background string

	exits []exit

	onSelector bool

	hotspots  []inkwell.Hotspot
	actors    []inkwell.SceneActor
	walkboxes []inkwell.Polygon
	onEnter   inkwell.Action
}

var screenDeck = []func() screen{
	screenPaulShop,
	screenNoodleHouse,
	screenAlley,
	screenNormanApartment,
	screenPoliceStation,
	screenHackerspace,
	screenIceCreamShop,
	screenTrinketShop,
	screenSmallRestaurant,
	screenSecretClub,
	screenBbsTerminal,
	screenStreet,
	screenHospital,
	screenServerFarm,
	screenSecretLab,
	screenRooftopHideout,
	screenPublicBBS,
	screenBvkBranch,
	screenCuratorShop,
	screenScrapMarket,
	screenSamizdatPress,
	screenShowroom,
	screenColumbarium,
}

func registerScreen(w *World) {
	screens := make([]screen, 0, len(screenDeck)+1)
	for _, f := range screenDeck {
		screens = append(screens, f())
	}
	for _, s := range append([]screen{screenSelector(screens)}, screens...) {
		w.G.SceneManager.Register(screenBuild(w, s))
	}
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

func screenBuild(w *World, s screen) inkwell.Scene {
	exits := s.exits
	if s.onSelector {
		exits = append(exits, toSelector())
	}
	hotspots := make([]inkwell.Hotspot, 0, len(exits)+len(s.hotspots))
	hotspots = append(hotspots, s.hotspots...)
	for _, e := range exits {
		hotspots = append(hotspots, e.hotspot(w))
	}

	actors := s.actors
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
	walkboxes := s.walkboxes
	if walkboxes == nil {
		walkboxes = screenFloor
	}
	onEnter := inkwell.Action(EnterScene(w, s.name))
	if s.onEnter != nil {
		onEnter = inkwell.Seq(onEnter, s.onEnter)
	}

	return inkwell.Scene{
		Name:       s.name,
		Title:      s.title,
		Background: s.background,
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
