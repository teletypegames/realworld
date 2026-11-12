package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/script"
)

var PaulDoor = SceneManager.Register(&Scene{
	Name: "paul_door",
	Title: Text{
		{Lang: "en", Text: "PAUL'S DOOR — A LETTER FROM NOODLE"},
		{Lang: "hu", Text: "PAUL AJTAJA — LEVÉL NOODLE-TŐL"},
	},
	Layers: []SceneLayer{
		{
			Asset: &Asset{
				Name: "paul_door",
				Path: "assets/bg/paul_door.png",
			},
			Z: -1,
		},
	},
	OnEnter: inkwell.Seq(
		inkwell.RunScript(script.PaulDoor),
		inkwell.GoTo(AircraftTakeoff),
	),
})
