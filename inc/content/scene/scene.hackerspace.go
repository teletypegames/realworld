package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name: inc.SceneHackerspace,
		Title: inkwell.Text{
			{Lang: "en", Text: "HACKERSPACE"},
			{Lang: "hu", Text: "HACKERSPACE"},
		},
		Background: inc.BgHackerspace,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
