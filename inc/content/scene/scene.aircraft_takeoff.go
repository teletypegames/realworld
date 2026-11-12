package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/script"
)

var AircraftTakeoff = SceneManager.Register(&Scene{
	Name: "aircraft_takeoff",
	Title: Text{
		{Lang: "en", Text: "AIRCRAFT — TAKE-OFF"},
		{Lang: "hu", Text: "REPÜLŐ — FELSZÁLLÁS"},
	},
	Layers: []SceneLayer{
		{
			Asset: &Asset{
				Name: "aircraft_takeoff",
				Path: "assets/bg/aircraft_takeoff.png",
			},
			Z: -1,
		},
	},
	Walkboxes: []Polygon{},
	Actors: []SceneActor{
		{
			Character: character.Aircraft,
			At: Point{
				X: -90,
				Y: 240,
			},
		},
	},
	OnEnter: inkwell.Seq(
		inkwell.RunScript(script.AircraftTakeoff),
		inkwell.GoTo(AircraftLanding),
	),
})
