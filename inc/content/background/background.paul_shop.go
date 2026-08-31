package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	BackgroundManager.Register(Background{
		Name: inc.BgPaulShop,
		Path: "assets/bg/paul_shop.png",
		Kind: inkwell.AssetImage,
	})
}
