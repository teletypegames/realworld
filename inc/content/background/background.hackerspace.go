package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgHackerspace,
		Path: "assets/bg/hackerspace.png",
		Kind: inkwell.AssetImage,
	})
}
