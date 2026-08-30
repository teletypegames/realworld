package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgColumbarium,
		Path: "assets/bg/columbarium.png",
		Kind: inkwell.AssetImage,
	})
}
