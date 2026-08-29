package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// scrapMarket — deck 20.
func backgroundScrapMarket() inkwell.Asset {
	return inkwell.Asset{
		Name: BgScrapMarket,
		Path: "assets/bg/scrap_market.png",
		Kind: inkwell.AssetImage,
	}
}
