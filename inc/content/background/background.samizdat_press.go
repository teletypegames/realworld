package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	BackgroundManager.Register(Background{
		Name: inc.BackgroundSamizdatPress,
		Path: "assets/bg/samizdat_press.png",
		Kind: inkwell.AssetImage,
	})
}
