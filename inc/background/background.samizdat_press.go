package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgSamizdatPress,
		Path: "assets/bg/samizdat_press.png",
		Kind: inkwell.AssetImage,
	})
}
