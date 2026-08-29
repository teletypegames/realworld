package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// trinketShop — deck 08.
func trinketShop() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgTrinketShop,
		Path: "assets/bg/trinket_shop.png",
		Kind: inkwell.AssetImage,
	}
}
