package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	SceneManager.Register(Scene{
		Name:       ScenePaulShop,
		Title:      "PAUL'S SHOP — JUNK AND GARAGE",
		Background: BgPaulShop,
		Exits: []inkwell.Exit{
			{
				To:    SceneNoodleHouse,
				Label: "the bus to San Francisco",
				Side:  inkwell.ExitNear,
			},
		},
	})
}
