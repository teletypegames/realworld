package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// noodleHouse — deck 02.
func backgroundNoodleHouse() inkwell.Asset {
	return inkwell.Asset{
		Name: BgNoodleHouse,
		Path: "assets/bg/noodle_house.png",
		Kind: inkwell.AssetImage,
	}
}
