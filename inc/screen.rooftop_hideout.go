package inc

// rooftopHideout is Norman's hiding place: an abandoned roof terrace with his
// personal things, his notes and an old BBS terminal — the emotional high point
// of the investigation (wiki: entities#helyszinek, timeline 23).
//
// It is up the stairs from his flat, and it is one of the places the catalogue
// puts a terminal.
func screenRooftopHideout() screen {
	return screen{
		name:       ScreenRooftopHideout,
		title:      "NORMAN'S ROOFTOP HIDEOUT",
		background: BgRooftopHideout,
		exits: []exit{
			{to: ScreenNormanApartment, label: "back down into the flat", side: sideNear},
			{to: ScreenBBSTerminal, label: "the old terminal", side: sideRight},
		},
	}
}
