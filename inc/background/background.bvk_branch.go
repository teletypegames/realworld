package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgBVKBranch,
		Path: "assets/bg/bvk_branch.png",
		Kind: inkwell.AssetImage,
	})
}
