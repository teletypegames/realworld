package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	BackgroundManager.Register(Background{
		Name: inc.BgAlley,
		Path: "assets/bg/alley.png",
		Kind: inkwell.AssetImage,
	})
}
