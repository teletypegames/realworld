package screen

import "git.teletypegames.org/games/realworld/internal/names"

// noodleHouse is the street front outside Noodle's place. By the time Paul
// gets there the police have sealed the house and taken Noodle in, and a
// uniform stands in front of the tape (wiki: story#noodle-szál).
//
// The wiki's graph runs NoodleUtcafront → Sikátor: the alley behind the house
// is reached from here, and from nowhere else.
func noodleHouse() screen {
	return screen{
		name:       names.ScreenNoodleHouse,
		title:      "NOODLE'S HOUSE — SEALED",
		background: names.BgNoodleHouse,
		onSelector: true,
		exits: []exit{
			{to: names.ScreenAlley, label: "the alley behind the house", side: sideRight},
		},
	}
}
