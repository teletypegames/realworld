package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	BackgroundManager.Register(Background{
		Name: BgIceCreamShop,
		Path: "assets/bg/ice_cream_shop.png",
		Kind: inkwell.AssetImage,
	})
}
