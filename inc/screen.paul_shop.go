package inc

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenPaulShop,
		Title:      "PAUL'S SHOP — JUNK AND GARAGE",
		Background: BgPaulShop,
		Exits: []exit{
			{
				to:    ScreenNoodleHouse,
				label: "the bus to San Francisco",
				side:  sideNear,
			},
		},
	})
}
