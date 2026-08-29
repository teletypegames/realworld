package inc

func screenSecretClub() screen {
	return screen{
		name:       ScreenSecretClub,
		title:      "SECRET CLUB — THE UNDERWATER SUN",
		background: BgSecretClub,
		exits: []exit{
			{to: ScreenBBSTerminal, label: "the terminal in the corner", side: sideBack},
			{to: ScreenTrinketShop, label: "back out through the shop", side: sideNear},
		},
	}
}
