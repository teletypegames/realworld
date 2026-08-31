package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgColumbarium,
		Path: "assets/bg/columbarium.png",
		Kind: inkwell.AssetImage,
	})
}
