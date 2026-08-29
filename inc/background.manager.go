package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

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
