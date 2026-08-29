package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// normanApartment — deck 03.
func normanApartment() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgNormanApartment,
		Path: "assets/bg/norman_apartment.png",
		Kind: inkwell.AssetImage,
	}
}
