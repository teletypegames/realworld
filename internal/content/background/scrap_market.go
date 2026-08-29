package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// scrapMarket — deck 20.
func scrapMarket() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgScrapMarket,
		Path: "assets/bg/scrap_market.png",
		Kind: inkwell.AssetImage,
	}
}
