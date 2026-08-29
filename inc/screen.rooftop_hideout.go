package inc

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
