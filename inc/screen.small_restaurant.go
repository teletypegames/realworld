package inc

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenSmallRestaurant,
		Title:      "SMALL RESTAURANT — NEXT DOOR",
		Background: BgSmallRestaurant,
		Exits: []exit{
			{
				to:    ScreenTrinketShop,
				label: "back to the trinket shop",
				side:  sideLeft,
			},
		},
	})
}
