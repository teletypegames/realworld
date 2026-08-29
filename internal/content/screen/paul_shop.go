package screen

import "git.teletypegames.org/games/realworld/internal/names"

// paulShop is where the game starts: the junk shop and garage Paul lives
// behind (wiki: entities#helyszinek). It is the only screen outside San
// Francisco — Noodle's letter arrives here and puts him on a bus
// (wiki: entities, Paul's timeline 1–2).
//
// The bus is a screen the deck does not have, so the exit is the journey
// itself: one way, with no coming back to the shop mid-story. That is also why
// the shop is not on the selector — it is not one of the city's locations.
func paulShop() screen {
	return screen{
		name:       names.ScreenPaulShop,
		title:      "PAUL'S SHOP — JUNK AND GARAGE",
		background: names.BgPaulShop,
		exits: []exit{
			{to: names.ScreenNoodleHouse, label: "the bus to San Francisco", side: sideNear},
		},
	}
}
