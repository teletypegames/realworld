package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.SceneShowroom,
		Title:      "NEUMATRONIC SHOWROOM",
		Background: inc.BgShowroom,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
