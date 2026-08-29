package inc

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenTrinketShop,
		Title:      "CHINATOWN — TRINKET SHOP",
		Background: BgTrinketShop,
		OnSelector: true,
		Exits: []exit{
			{
				to:    ScreenSmallRestaurant,
				label: "the eating place next door",
				side:  sideRight,
			},
			{
				to:      ScreenSecretClub,
				label:   "the way in at the back",
				side:    sideBack,
				needs:   FlagClubEntry,
				blocked: "Just a shop, as far as the man behind the counter is concerned.",
			},
		},
	})
}
