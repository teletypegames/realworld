package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	BackgroundManager.Register(Background{
		Name: inc.BackgroundIceCreamShop,
		Path: "assets/bg/ice_cream_shop.png",
		Kind: inkwell.AssetImage,
	})
}
