package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgNoodleHouse,
		Path: "assets/bg/noodle_house.png",
		Kind: inkwell.AssetImage,
	})
}
