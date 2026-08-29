package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	SceneManager.Register(Scene{
		Name:       SceneNormanApartment,
		Title:      "NORMAN'S APARTMENT",
		Background: BgNormanApartment,
		Exits: []inkwell.Exit{
			exitToSelector,
			{
				To:    SceneRooftopHideout,
				Label: "the stairs up to the roof",
				Side:  inkwell.ExitBack,
			},
			{
				To:    SceneBBSTerminal,
				Label: "Norman's terminal",
				Side:  inkwell.ExitRight,
			},
		},
	})
}
