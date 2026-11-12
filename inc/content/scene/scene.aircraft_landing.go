package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/script"
)

var AircraftLanding = SceneManager.Register(&Scene{
	Name: "aircraft_landing",
	Title: Text{
		{Lang: "en", Text: "AIRCRAFT — LANDING IN SAN FRANCISCO"},
		{Lang: "hu", Text: "REPÜLŐ — LESZÁLLÁS SAN FRANCISCÓBAN"},
	},
	Layers: []SceneLayer{
		{
			Asset: &Asset{
				Name: "aircraft_landing",
				Path: "assets/bg/aircraft_landing.png",
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
				Y: 70,
			},
		},
	},
	OnEnter: inkwell.Seq(
		inkwell.RunScript(script.AircraftLanding),
		inkwell.GoTo(Bus),
	),
})
