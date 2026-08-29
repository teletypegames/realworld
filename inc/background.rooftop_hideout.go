package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	BackgroundManager.Register(Background{
		Name: BgRooftopHideout,
		Path: "assets/bg/rooftop_hideout.png",
		Kind: inkwell.AssetImage,
	})
}
