package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgRooftopHideout,
		Path: "assets/bg/rooftop_hideout.png",
		Kind: inkwell.AssetImage,
	})
}
