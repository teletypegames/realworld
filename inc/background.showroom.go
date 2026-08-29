package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	BackgroundManager.Register(Background{
		Name: BgShowroom,
		Path: "assets/bg/showroom.png",
		Kind: inkwell.AssetImage,
	})
}
