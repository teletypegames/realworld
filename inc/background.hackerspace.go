package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	BackgroundManager.Register(Background{
		Name: BgHackerspace,
		Path: "assets/bg/hackerspace.png",
		Kind: inkwell.AssetImage,
	})
}
