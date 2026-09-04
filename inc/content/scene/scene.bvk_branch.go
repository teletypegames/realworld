package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

var BVKBranch = SceneManager.Register(&Scene{
	Name: "bvk_branch",
	Title: inkwell.Text{
		{Lang: "en", Text: "BVK — SAN FRANCISCO BRANCH"},
		{Lang: "hu", Text: "BVK — SAN FRANCISCÓI FIÓK"},
	},
})

var _ = BVKBranch.Leads(
	exitToSelector,
)
