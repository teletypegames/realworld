package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	SceneManager.Register(Scene{
		Name:       SceneRooftopHideout,
		Title:      "NORMAN'S ROOFTOP HIDEOUT",
		Background: BgRooftopHideout,
		Exits: []inkwell.Exit{
			{
				To:    SceneNormanApartment,
				Label: "back down into the flat",
				Side:  inkwell.ExitNear,
			},
			{
				To:    SceneBBSTerminal,
				Label: "the old terminal",
				Side:  inkwell.ExitRight,
			},
		},
	})
}
