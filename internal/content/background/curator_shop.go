package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// curatorShop — deck 19.
func curatorShop() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgCuratorShop,
		Path: "assets/bg/curator_shop.png",
		Kind: inkwell.AssetImage,
	}
}
