package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.SceneBVKBranch,
		Title:      "BVK — SAN FRANCISCO BRANCH",
		Background: inc.BgBVKBranch,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
