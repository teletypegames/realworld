package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// hospital — deck 13.
func hospital() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgHospital,
		Path: "assets/bg/hospital.png",
		Kind: inkwell.AssetImage,
	}
}
