package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	BackgroundManager.Register(Background{
		Name: inc.BgPublicBBS,
		Path: "assets/bg/public_bbs.png",
		Kind: inkwell.AssetImage,
	})
}
