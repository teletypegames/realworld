package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	SceneManager.Register(Scene{
		Name:       SceneColumbarium,
		Title:      "COLUMBARIUM — DEX'S MEMORIAL",
		Background: BgColumbarium,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
