package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type bvkBranch struct{}

var BVKBranch = SceneManager.Register(&Scene{
	Name: "bvk_branch",
	Title: inkwell.Text{
		{Lang: "en", Text: "BVK — SAN FRANCISCO BRANCH"},
		{Lang: "hu", Text: "BVK — SAN FRANCISCÓI FIÓK"},
	},
	Exits: []inkwell.Exit{
		exitToSelector,
	},
})

func (bvkBranch) Scene() *inkwell.Scene { return BVKBranch }
