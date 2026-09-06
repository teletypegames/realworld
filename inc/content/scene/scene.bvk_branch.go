package scene

import ()

var BVKBranch = SceneManager.Register(&Scene{
	Name: "bvk_branch",
	Title: Text{
		{Lang: "en", Text: "BVK — SAN FRANCISCO BRANCH"},
		{Lang: "hu", Text: "BVK — SAN FRANCISCÓI FIÓK"},
	},
	Layers: []SceneLayer{
		{
			Asset: &Asset{
				Name: "bvk_branch",
				Path: "assets/bg/bvk_branch.png",
			},
			Z: -1,
		},
	},
})

var _ = BVKBranch.Leads(
	exitToSelector,
)
