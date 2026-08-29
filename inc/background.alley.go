package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func backgroundAlley() inkwell.Asset {
	return inkwell.Asset{
		Name: BgAlley,
		Path: "assets/bg/alley.png",
		Kind: inkwell.AssetImage,
	}
}
