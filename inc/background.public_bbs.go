package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func backgroundPublicBBS() inkwell.Asset {
	return inkwell.Asset{
		Name: BgPublicBBS,
		Path: "assets/bg/public_bbs.png",
		Kind: inkwell.AssetImage,
	}
}
