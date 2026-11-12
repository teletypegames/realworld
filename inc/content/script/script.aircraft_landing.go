package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
)

var AircraftLanding = ScriptManager.Register(&Script{
	Name: "aircraft_landing",
	Actions: inkwell.Seq(
		inkwell.Wait(0.5),
		inkwell.Walk(character.Aircraft, inkwell.Point{
			X: 420,
			Y: 230,
		}),
		inkwell.Walk(character.Aircraft, inkwell.Point{
			X: 760,
			Y: 250,
		}),
		inkwell.Wait(0.5),
	),
})
