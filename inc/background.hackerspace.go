package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func backgroundHackerspace() inkwell.Asset {
	return inkwell.Asset{
		Name: BgHackerspace,
		Path: "assets/bg/hackerspace.png",
		Kind: inkwell.AssetImage,
	}
}
