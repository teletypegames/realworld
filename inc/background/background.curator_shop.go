package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgCuratorShop,
		Path: "assets/bg/curator_shop.png",
		Kind: inkwell.AssetImage,
	})
}
