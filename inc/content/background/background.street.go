package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	BackgroundManager.Register(Background{
		Name: inc.BackgroundStreet,
		Path: "assets/bg/street.png",
		Kind: inkwell.AssetImage,
	})
}
