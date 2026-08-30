package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgIceCreamShop,
		Path: "assets/bg/ice_cream_shop.png",
		Kind: inkwell.AssetImage,
	})
}
