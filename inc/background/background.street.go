package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgStreet,
		Path: "assets/bg/street.png",
		Kind: inkwell.AssetImage,
	})
}
