package inc

func screenNoodleHouse() screen {
	return screen{
		name:       ScreenNoodleHouse,
		title:      "NOODLE'S HOUSE — SEALED",
		background: BgNoodleHouse,
		onSelector: true,
		exits: []exit{
			{to: ScreenAlley, label: "the alley behind the house", side: sideRight},
		},
	}
}
