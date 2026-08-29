package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	BackgroundManager.Register(Background{
		Name: BgPublicBBS,
		Path: "assets/bg/public_bbs.png",
		Kind: inkwell.AssetImage,
	})
}
