package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.ScenePaulShop,
		Title:      "PAUL'S SHOP — JUNK AND GARAGE",
		Background: inc.BgPaulShop,
		Exits: []inkwell.Exit{
			{
				To:    inc.SceneNoodleHouse,
				Label: "the bus to San Francisco",
				Side:  inkwell.ExitNear,
			},
		},
	})
}
