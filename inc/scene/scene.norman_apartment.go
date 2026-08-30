package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.SceneNormanApartment,
		Title:      "NORMAN'S APARTMENT",
		Background: inc.BgNormanApartment,
		Exits: []inkwell.Exit{
			exitToSelector,
			{
				To:    inc.SceneRooftopHideout,
				Label: "the stairs up to the roof",
				Side:  inkwell.ExitBack,
			},
			{
				To:    inc.SceneBBSTerminal,
				Label: "Norman's terminal",
				Side:  inkwell.ExitRight,
			},
		},
	})
}
