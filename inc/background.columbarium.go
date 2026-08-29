package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	BackgroundManager.Register(Background{
		Name: BgColumbarium,
		Path: "assets/bg/columbarium.png",
		Kind: inkwell.AssetImage,
	})
}
