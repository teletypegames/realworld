package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
)

var AircraftTakeoff = ScriptManager.Register(&Script{
	Name: "aircraft_takeoff",
	Actions: inkwell.Seq(
		inkwell.Wait(0.5),
		inkwell.Walk(character.Aircraft, inkwell.Point{
			X: 220,
			Y: 238,
		}),
		inkwell.Walk(character.Aircraft, inkwell.Point{
			X: 760,
			Y: -20,
		}),
		inkwell.Wait(0.5),
	),
})
