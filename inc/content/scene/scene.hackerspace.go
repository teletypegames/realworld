package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	SceneManager.Register(Scene{
		Name: inc.SceneHackerspace,
		Title: inkwell.Text{
			{Lang: "en", Text: "HACKERSPACE"},
			{Lang: "hu", Text: "HACKERSPACE"},
		},
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
