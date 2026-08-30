package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.SceneServerFarm,
		Title:      "SERVER FARM",
		Background: inc.BgServerFarm,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
