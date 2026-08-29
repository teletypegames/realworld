package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func backgroundPaulShop() inkwell.Asset {
	return inkwell.Asset{
		Name: BgPaulShop,
		Path: "assets/bg/paul_shop.png",
		Kind: inkwell.AssetImage,
	}
}
