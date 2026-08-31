package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	BackgroundManager.Register(Background{
		Name: inc.BgRooftopHideout,
		Path: "assets/bg/rooftop_hideout.png",
		Kind: inkwell.AssetImage,
	})
}
