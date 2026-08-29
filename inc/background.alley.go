package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	BackgroundManager.Register(Background{
		Name: BgAlley,
		Path: "assets/bg/alley.png",
		Kind: inkwell.AssetImage,
	})
}
