package inc

func screenSmallRestaurant() screen {
	return screen{
		name:       ScreenSmallRestaurant,
		title:      "SMALL RESTAURANT — NEXT DOOR",
		background: BgSmallRestaurant,
		exits: []exit{
			{to: ScreenTrinketShop, label: "back to the trinket shop", side: sideLeft},
		},
	}
}
