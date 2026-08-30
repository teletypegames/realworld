package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgPaulShop,
		Path: "assets/bg/paul_shop.png",
		Kind: inkwell.AssetImage,
	})
}
