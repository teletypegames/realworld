package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	BackgroundManager.Register(Background{
		Name: BgPaulShop,
		Path: "assets/bg/paul_shop.png",
		Kind: inkwell.AssetImage,
	})
}
