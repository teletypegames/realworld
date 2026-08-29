package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// iceCreamShop — deck 07.
func iceCreamShop() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgIceCreamShop,
		Path: "assets/bg/ice_cream_shop.png",
		Kind: inkwell.AssetImage,
	}
}
