package inc

func screenTrinketShop() screen {
	return screen{
		name:       ScreenTrinketShop,
		title:      "CHINATOWN — TRINKET SHOP",
		background: BgTrinketShop,
		onSelector: true,
		exits: []exit{
			{to: ScreenSmallRestaurant, label: "the eating place next door", side: sideRight},
			{
				to: ScreenSecretClub, label: "the way in at the back", side: sideBack,
				needs:   FlagClubEntry,
				blocked: "Just a shop, as far as the man behind the counter is concerned.",
			},
		},
	}
}
