package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name: inc.SceneSecretLab,
		Title: inkwell.Text{
			{Lang: "en", Text: "SECRET RESEARCH LABORATORY"},
			{Lang: "hu", Text: "TITKOS KUTATÓLABOR"},
		},
		Background: inc.BgSecretLab,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
