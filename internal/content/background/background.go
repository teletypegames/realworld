// Package background registers one image asset per screen — one file each.
//
// A background is the whole of a screen's art: inkwell scales it over the
// entire window, so every file here points at a 640×380 PNG under assets/bg/,
// the size the concept-art deck is drawn at. Two of them are not painted yet
// and fall back to inkwell's placeholder.
//
// The order below is the deck's, which is the wiki's location-catalogue order
// (projects/realworld/entities#helyszinek).
package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/world"
)

// Register adds every background to the game.
func Register(w *world.World) {
	for _, a := range []inkwell.Asset{
		selector(),
		paulShop(),
		noodleHouse(),
		normanApartment(),
		policeStation(),
		alley(),
		hackerspace(),
		iceCreamShop(),
		trinketShop(),
		smallRestaurant(),
		secretClub(),
		bbsTerminal(),
		street(),
		hospital(),
		serverFarm(),
		secretLab(),
		rooftopHideout(),
		publicBbs(),
		bvkBranch(),
		curatorShop(),
		scrapMarket(),
		samizdatPress(),
		showroom(),
		columbarium(),
	} {
		w.G.AssetManager.Register(a)
	}
}
