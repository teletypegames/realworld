package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	BackgroundManager.Register(Background{
		Name: BgBBSTerminal,
		Path: "assets/bg/bbs_terminal.png",
		Kind: inkwell.AssetImage,
	})
}
