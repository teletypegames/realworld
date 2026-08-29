package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// bvkBranch — deck 18.
func bvkBranch() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgBVKBranch,
		Path: "assets/bg/bvk_branch.png",
		Kind: inkwell.AssetImage,
	}
}
