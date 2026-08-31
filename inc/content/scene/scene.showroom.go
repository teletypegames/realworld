package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	SceneManager.Register(Scene{
		Name: inc.SceneShowroom,
		Title: inkwell.Text{
			{Lang: "en", Text: "NEUMATRONIC SHOWROOM"},
			{Lang: "hu", Text: "NEUMATRONIC BEMUTATÓTEREM"},
		},
		Background: inc.BgShowroom,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
