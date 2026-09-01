package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	SceneManager.Register(Scene{
		Name: inc.SceneSecretLab,
		Title: inkwell.Text{
			{Lang: "en", Text: "SECRET RESEARCH LABORATORY"},
			{Lang: "hu", Text: "TITKOS KUTATÓLABOR"},
		},
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
