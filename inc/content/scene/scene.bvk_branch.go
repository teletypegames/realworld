package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	SceneManager.Register(Scene{
		Name: inc.SceneBVKBranch,
		Title: inkwell.Text{
			{Lang: "en", Text: "BVK — SAN FRANCISCO BRANCH"},
			{Lang: "hu", Text: "BVK — SAN FRANCISCÓI FIÓK"},
		},
		Background: inc.BackgroundBVKBranch,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
