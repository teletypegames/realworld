package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgTrinketShop,
		Path: "assets/bg/trinket_shop.png",
		Kind: inkwell.AssetImage,
	})
}
