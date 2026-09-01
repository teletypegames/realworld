package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	SceneManager.Register(Scene{
		Name: inc.SceneColumbarium,
		Title: inkwell.Text{
			{Lang: "en", Text: "COLUMBARIUM — DEX'S MEMORIAL"},
			{Lang: "hu", Text: "KOLUMBÁRIUM — DEX EMLÉKHELYE"},
		},
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
