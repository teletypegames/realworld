package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	SceneManager.Register(Scene{
		Name:       SceneCuratorShop,
		Title:      "CURATOR SERVICE — THE MASTER'S WORKSHOP",
		Background: BgCuratorShop,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
