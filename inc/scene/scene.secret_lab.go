package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.SceneSecretLab,
		Title:      "SECRET RESEARCH LABORATORY",
		Background: inc.BgSecretLab,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
