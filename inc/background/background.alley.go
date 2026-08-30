package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgAlley,
		Path: "assets/bg/alley.png",
		Kind: inkwell.AssetImage,
	})
}
