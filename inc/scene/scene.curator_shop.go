package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.SceneCuratorShop,
		Title:      "CURATOR SERVICE — THE MASTER'S WORKSHOP",
		Background: inc.BgCuratorShop,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
