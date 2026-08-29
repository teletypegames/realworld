package inc

// normanApartment is the flat Norman was taken from: his Armilla left behind
// on the side, his Personal Tape missing from it, his terminal full of the
// correspondence and the diary (wiki: story#norman-lakása).
//
// Two ways on from here. His terminal is one of the BBS access points the
// catalogue lists ("a titkos klub ECHO-terminálja; később Norman lakása /
// tetőterasz"), and the roof hideout is up his own stairs.
func screenNormanApartment() screen {
	return screen{
		name:       ScreenNormanApartment,
		title:      "NORMAN'S APARTMENT",
		background: BgNormanApartment,
		onSelector: true,
		exits: []exit{
			{to: ScreenRooftopHideout, label: "the stairs up to the roof", side: sideBack},
			{to: ScreenBBSTerminal, label: "Norman's terminal", side: sideRight},
		},
	}
}
