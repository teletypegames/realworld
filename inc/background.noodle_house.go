package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	BackgroundManager.Register(Background{
		Name: BgNoodleHouse,
		Path: "assets/bg/noodle_house.png",
		Kind: inkwell.AssetImage,
	})
}
