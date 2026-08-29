package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	BackgroundManager.Register(Background{
		Name: BgSamizdatPress,
		Path: "assets/bg/samizdat_press.png",
		Kind: inkwell.AssetImage,
	})
}
