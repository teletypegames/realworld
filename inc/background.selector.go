package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	BackgroundManager.Register(Background{
		Name: BgSelector,
		Path: "assets/bg/selector.png",
		Kind: inkwell.AssetImage,
	})
}
