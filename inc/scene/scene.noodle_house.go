package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.SceneNoodleHouse,
		Title:      "NOODLE'S HOUSE — SEALED",
		Background: inc.BgNoodleHouse,
		Exits: []inkwell.Exit{
			exitToSelector,
			{
				To:    inc.SceneAlley,
				Label: "the alley behind the house",
				Side:  inkwell.ExitRight,
			},
		},
	})
}
