package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	BackgroundManager.Register(Background{
		Name: BgScrapMarket,
		Path: "assets/bg/scrap_market.png",
		Kind: inkwell.AssetImage,
	})
}
