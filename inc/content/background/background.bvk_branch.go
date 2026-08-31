package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	BackgroundManager.Register(Background{
		Name: inc.BgBVKBranch,
		Path: "assets/bg/bvk_branch.png",
		Kind: inkwell.AssetImage,
	})
}
