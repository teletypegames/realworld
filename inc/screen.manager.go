package inc

// Screens — one file per screen.
//
// "Screen" is this project's word: the wiki's catalogue, the concept-art deck
// and the beat tables all count screens. inkwell's word for the same thing is
// Scene, which is the type every file here returns — a screen IS the engine's
// scene, and the rest of the code says "scene" only where it is talking to the
// engine.
//
// A file here holds everything that is true of one screen: its picture, the
// name on the top bar, where you can get to from it (screen.exit.go), and — once its
// beat is written — its hotspots, NPCs and dialogue. Only the alley has that
// last part so far; the rest are greybox.

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// screen is what each file in this package fills in.
type screen struct {
	name       string
	title      string
	background string

	// exits: where you can get to from here. See screen.exit.go.
	exits []exit
	// onSelector marks a main location — one of those the wiki's
	// Helyszínválasztó lists (story#helyszín-gráf-vázlat). It gets a way back
	// to the selector for free, and the selector gets a way here.
	onSelector bool

	// What an authored beat adds on top. Zero values are fine: Paul lands in
	// the middle of the default floor strip and nothing else happens.
	hotspots  []inkwell.Hotspot
	actors    []inkwell.SceneActor
	walkboxes []inkwell.Polygon
	onEnter   inkwell.Action
}

// screenDeck is every screen in the game, in the order of the wiki's location
// catalogue (entities#helyszinek) — which the concept-art deck follows number
// for number. Registration order is this order, and it is also the order the
// arrow keys walk (ui.screen_nav.go).
//
// Deck 05, Norman's workplace, is in the catalogue but has not been drawn, so
// it has no file yet. The alley has no deck number: the catalogue keeps it in
// the same row as Noodle's house, and that is where it sits here too.
var screenDeck = []func() screen{
	screenPaulShop,        // 01
	screenNoodleHouse,     // 02
	screenAlley,           // 02, same catalogue row
	screenNormanApartment, // 03
	screenPoliceStation,   // 04
	screenHackerspace,     // 06
	screenIceCreamShop,    // 07
	screenTrinketShop,     // 08
	screenSmallRestaurant, // 09
	screenSecretClub,      // 10
	screenBbsTerminal,     // 11
	screenStreet,          // 12
	screenHospital,        // 13
	screenServerFarm,      // 14
	screenSecretLab,       // 15
	screenRooftopHideout,  // 16
	screenPublicBBS,       // 17
	screenBvkBranch,       // 18
	screenCuratorShop,     // 19
	screenScrapMarket,     // 20
	screenSamizdatPress,   // 21
	screenShowroom,        // 22
	screenColumbarium,     // 23
}

// registerScreen adds every screen to the game, the selector first: it is the centre
// of the wiki's location graph, and it is built from the screenDeck rather than
// listed in it.
func registerScreen(w *World) {
	screens := make([]screen, 0, len(screenDeck)+1)
	for _, f := range screenDeck {
		screens = append(screens, f())
	}
	for _, s := range append([]screen{screenSelector(screens)}, screens...) {
		w.G.SceneManager.Register(screenBuild(w, s))
	}
}

// The default floor: the bottom of the scene region, above the HUD divider.
// A screen with a painted floor of its own overrides it.
var screenFloor = []inkwell.Polygon{
	inkwell.Poly(
		inkwell.Point{X: 16, Y: 196}, inkwell.Point{X: 624, Y: 196},
		inkwell.Point{X: 624, Y: 258}, inkwell.Point{X: 16, Y: 258},
	),
}

// screenBuild turns a screen into the engine's scene.
//
// The authored hotspots go in front of the exits, because the engine takes the
// first one whose area contains the click (HotspotAt): a painted thing beats
// the edge strip an exit sits on, wherever the two overlap.
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
			{CharacterName: Paul, At: inkwell.Point{X: 320, Y: 232}},
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

// setVarIfEmpty only writes when the variable is still unset, so re-entering a
// screen does not clobber what the game has set in the meantime.
func setVarIfEmpty(name string, v any) inkwell.Action {
	return Fn(func(ctx *inkwell.Ctx) {
		if ctx.Game.State.Var(name) == nil {
			ctx.Game.State.SetVar(name, v)
		}
	})
}
