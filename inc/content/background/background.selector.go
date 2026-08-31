package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	BackgroundManager.Register(Background{
		Name: inc.BgSelector,
		Path: "assets/bg/selector.png",
		Kind: inkwell.AssetImage,
	})
}
