package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	BackgroundManager.Register(Background{
		Name: inc.BackgroundColumbarium,
		Path: "assets/bg/columbarium.png",
		Kind: inkwell.AssetImage,
	})
}
