package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	BackgroundManager.Register(Background{
		Name: BgCuratorShop,
		Path: "assets/bg/curator_shop.png",
		Kind: inkwell.AssetImage,
	})
}
