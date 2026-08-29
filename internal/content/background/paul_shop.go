package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// paulShop — deck 01.
func paulShop() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgPaulShop,
		Path: "assets/bg/paul_shop.png",
		Kind: inkwell.AssetImage,
	}
}
