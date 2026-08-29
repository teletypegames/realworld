package inc

// trinketShop is the Chinatown shop at the club's address, whose owner sells
// Paul junk he does not need and is, in fact, the club's doorman
// (wiki: story#csecsebecse-bolt--kajálda-szál).
//
// Both of the wiki's arrows out of it are here: Bolt → Kajálda, the eating
// place next door, and Bolt →|beengedés| Klub, which stays shut until the
// shopkeeper has been shown the underwater sun.
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
