package inc

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenSecretClub,
		Title:      "SECRET CLUB — THE UNDERWATER SUN",
		Background: BgSecretClub,
		Exits: []exit{
			{
				to:    ScreenBBSTerminal,
				label: "the terminal in the corner",
				side:  sideBack,
			},
			{
				to:    ScreenTrinketShop,
				label: "back out through the shop",
				side:  sideNear,
			},
		},
	})
}
