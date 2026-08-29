package inc

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenNoodleHouse,
		Title:      "NOODLE'S HOUSE — SEALED",
		Background: BgNoodleHouse,
		OnSelector: true,
		Exits: []exit{
			{
				to:    ScreenAlley,
				label: "the alley behind the house",
				side:  sideRight,
			},
		},
	})
}
