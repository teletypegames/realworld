package inc

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
