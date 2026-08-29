package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// smallRestaurant — deck 09.
func backgroundSmallRestaurant() inkwell.Asset {
	return inkwell.Asset{
		Name: BgSmallRestaurant,
		Path: "assets/bg/small_restaurant.png",
		Kind: inkwell.AssetImage,
	}
}
