package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// showroom — deck 22.
func backgroundShowroom() inkwell.Asset {
	return inkwell.Asset{
		Name: BgShowroom,
		Path: "assets/bg/showroom.png",
		Kind: inkwell.AssetImage,
	}
}
