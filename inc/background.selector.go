package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func backgroundSelector() inkwell.Asset {
	return inkwell.Asset{
		Name: BgSelector,
		Path: "assets/bg/selector.png",
		Kind: inkwell.AssetImage,
	}
}
