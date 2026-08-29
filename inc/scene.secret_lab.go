package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	SceneManager.Register(Scene{
		Name:       SceneSecretLab,
		Title:      "SECRET RESEARCH LABORATORY",
		Background: BgSecretLab,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
