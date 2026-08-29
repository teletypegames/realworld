package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	SceneManager.Register(Scene{
		Name:       SceneBVKBranch,
		Title:      "BVK — SAN FRANCISCO BRANCH",
		Background: BgBVKBranch,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
