package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name: inc.SceneBVKBranch,
		Title: inkwell.Text{
			{Lang: "en", Text: "BVK — SAN FRANCISCO BRANCH"},
			{Lang: "hu", Text: "BVK — SAN FRANCISCÓI FIÓK"},
		},
		Background: inc.BgBVKBranch,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
