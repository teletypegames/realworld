package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgSmallRestaurant,
		Path: "assets/bg/small_restaurant.png",
		Kind: inkwell.AssetImage,
	})
}
