package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	BackgroundManager.Register(Background{
		Name: BgTrinketShop,
		Path: "assets/bg/trinket_shop.png",
		Kind: inkwell.AssetImage,
	})
}
