package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	SceneManager.Register(Scene{
		Name:       SceneShowroom,
		Title:      "NEUMATRONIC SHOWROOM",
		Background: BgShowroom,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
