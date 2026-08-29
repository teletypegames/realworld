package screen

import "git.teletypegames.org/games/realworld/internal/names"

// smallRestaurant is the little Chinese place beside the trinket shop, and the
// lit mirror in its toilet is the puzzle: it flickers, and something shows
// behind it (wiki: story, the Kajálda beat).
//
// The wiki's graph runs Kajálda → étterem rész → WC; the deck paints all three
// as one screen, so the mirror is a hotspot here. The only way out is back to
// the shop — this is an inner room, not a location on the selector.
func smallRestaurant() screen {
	return screen{
		name:       names.ScreenSmallRestaurant,
		title:      "SMALL RESTAURANT — NEXT DOOR",
		background: names.BgSmallRestaurant,
		exits: []exit{
			{to: names.ScreenTrinketShop, label: "back to the trinket shop", side: sideLeft},
		},
	}
}
