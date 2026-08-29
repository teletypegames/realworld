package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	SceneManager.Register(Scene{
		Name:       SceneHospital,
		Title:      "HOSPITAL — PSYCHIATRIC WING",
		Background: BgHospital,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
