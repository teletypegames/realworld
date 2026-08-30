package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgPublicBBS,
		Path: "assets/bg/public_bbs.png",
		Kind: inkwell.AssetImage,
	})
}
