package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func backgroundSmallRestaurant() inkwell.Asset {
	return inkwell.Asset{
		Name: BgSmallRestaurant,
		Path: "assets/bg/small_restaurant.png",
		Kind: inkwell.AssetImage,
	}
}
