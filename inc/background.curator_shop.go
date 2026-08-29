package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func backgroundCuratorShop() inkwell.Asset {
	return inkwell.Asset{
		Name: BgCuratorShop,
		Path: "assets/bg/curator_shop.png",
		Kind: inkwell.AssetImage,
	}
}
