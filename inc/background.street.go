package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func backgroundStreet() inkwell.Asset {
	return inkwell.Asset{
		Name: BgStreet,
		Path: "assets/bg/street.png",
		Kind: inkwell.AssetImage,
	}
}
