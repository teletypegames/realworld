package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	BackgroundManager.Register(Background{
		Name: inc.BackgroundAlley,
		Path: "assets/bg/alley.png",
		Kind: inkwell.AssetImage,
	})
}
