package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgShowroom,
		Path: "assets/bg/showroom.png",
		Kind: inkwell.AssetImage,
	})
}
