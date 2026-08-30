package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgScrapMarket,
		Path: "assets/bg/scrap_market.png",
		Kind: inkwell.AssetImage,
	})
}
