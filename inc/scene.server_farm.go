package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	SceneManager.Register(Scene{
		Name:       SceneServerFarm,
		Title:      "SERVER FARM",
		Background: BgServerFarm,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
