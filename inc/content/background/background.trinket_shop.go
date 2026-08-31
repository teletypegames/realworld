package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	BackgroundManager.Register(Background{
		Name: inc.BgTrinketShop,
		Path: "assets/bg/trinket_shop.png",
		Kind: inkwell.AssetImage,
	})
}
