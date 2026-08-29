package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	BackgroundManager.Register(Background{
		Name: BgBVKBranch,
		Path: "assets/bg/bvk_branch.png",
		Kind: inkwell.AssetImage,
	})
}
