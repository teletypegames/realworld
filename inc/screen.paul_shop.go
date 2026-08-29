package inc

func screenPaulShop() screen {
	return screen{
		name:       ScreenPaulShop,
		title:      "PAUL'S SHOP — JUNK AND GARAGE",
		background: BgPaulShop,
		exits: []exit{
			{
				to:    ScreenNoodleHouse,
				label: "the bus to San Francisco",
				side:  sideNear,
			},
		},
	}
}
