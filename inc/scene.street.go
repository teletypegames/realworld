package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	SceneManager.Register(Scene{
		Name:       SceneStreet,
		Title:      "STREET",
		Background: BgStreet,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
