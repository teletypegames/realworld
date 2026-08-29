package inc

// Backgrounds — one image asset per screen, one file each.
//
// A background is the whole of a screen's art: inkwell scales it over the
// entire window, so every file here points at a 640×380 PNG under assets/bg/,
// the size the concept-art deck is drawn at. Two of them are not painted yet
// and fall back to inkwell's placeholder.
//
// The order below is the deck's, which is the wiki's location-catalogue order
// (projects/realworld/entities#helyszinek).

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// registerBackground adds every background to the game.
func registerBackground(w *World) {
	for _, a := range []inkwell.Asset{
		backgroundSelector(),
		backgroundPaulShop(),
		backgroundNoodleHouse(),
		backgroundNormanApartment(),
		backgroundPoliceStation(),
		backgroundAlley(),
		backgroundHackerspace(),
		backgroundIceCreamShop(),
		backgroundTrinketShop(),
		backgroundSmallRestaurant(),
		backgroundSecretClub(),
		backgroundBbsTerminal(),
		backgroundStreet(),
		backgroundHospital(),
		backgroundServerFarm(),
		backgroundSecretLab(),
		backgroundRooftopHideout(),
		backgroundPublicBBS(),
		backgroundBvkBranch(),
		backgroundCuratorShop(),
		backgroundScrapMarket(),
		backgroundSamizdatPress(),
		backgroundShowroom(),
		backgroundColumbarium(),
	} {
		w.G.AssetManager.Register(a)
	}
}
