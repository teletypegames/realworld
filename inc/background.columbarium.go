package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// columbarium — deck 23.
func backgroundColumbarium() inkwell.Asset {
	return inkwell.Asset{
		Name: BgColumbarium,
		Path: "assets/bg/columbarium.png",
		Kind: inkwell.AssetImage,
	}
}
