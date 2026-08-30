package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.SceneColumbarium,
		Title:      "COLUMBARIUM — DEX'S MEMORIAL",
		Background: inc.BgColumbarium,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
