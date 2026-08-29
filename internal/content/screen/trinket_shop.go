package screen

import "git.teletypegames.org/games/realworld/internal/names"

// trinketShop is the Chinatown shop at the club's address, whose owner sells
// Paul junk he does not need and is, in fact, the club's doorman
// (wiki: story#csecsebecse-bolt--kajálda-szál).
//
// Both of the wiki's arrows out of it are here: Bolt → Kajálda, the eating
// place next door, and Bolt →|beengedés| Klub, which stays shut until the
// shopkeeper has been shown the underwater sun.
func trinketShop() screen {
	return screen{
		name:       names.ScreenTrinketShop,
		title:      "CHINATOWN — TRINKET SHOP",
		background: names.BgTrinketShop,
		onSelector: true,
		exits: []exit{
			{to: names.ScreenSmallRestaurant, label: "the eating place next door", side: sideRight},
			{
				to: names.ScreenSecretClub, label: "the way in at the back", side: sideBack,
				needs:   names.FlagClubEntry,
				blocked: "Just a shop, as far as the man behind the counter is concerned.",
			},
		},
	}
}
