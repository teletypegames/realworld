package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// showroom — deck 22.
func showroom() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgShowroom,
		Path: "assets/bg/showroom.png",
		Kind: inkwell.AssetImage,
	}
}
