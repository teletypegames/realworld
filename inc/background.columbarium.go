package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func backgroundColumbarium() inkwell.Asset {
	return inkwell.Asset{
		Name: BgColumbarium,
		Path: "assets/bg/columbarium.png",
		Kind: inkwell.AssetImage,
	}
}
