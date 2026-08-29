package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	SceneManager.Register(Scene{
		Name:       SceneNoodleHouse,
		Title:      "NOODLE'S HOUSE — SEALED",
		Background: BgNoodleHouse,
		Exits: []inkwell.Exit{
			exitToSelector,
			{
				To:    SceneAlley,
				Label: "the alley behind the house",
				Side:  inkwell.ExitRight,
			},
		},
	})
}
