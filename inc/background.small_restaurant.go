package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	BackgroundManager.Register(Background{
		Name: BgSmallRestaurant,
		Path: "assets/bg/small_restaurant.png",
		Kind: inkwell.AssetImage,
	})
}
