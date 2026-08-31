package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgBBSTerminal,
		Path: "assets/bg/bbs_terminal.png",
		Kind: inkwell.AssetImage,
	})
}
