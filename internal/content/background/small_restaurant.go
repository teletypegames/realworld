package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// smallRestaurant — deck 09.
func smallRestaurant() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgSmallRestaurant,
		Path: "assets/bg/small_restaurant.png",
		Kind: inkwell.AssetImage,
	}
}
