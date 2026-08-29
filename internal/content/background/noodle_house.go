package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// noodleHouse — deck 02.
func noodleHouse() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgNoodleHouse,
		Path: "assets/bg/noodle_house.png",
		Kind: inkwell.AssetImage,
	}
}
