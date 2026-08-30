package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.SceneRooftopHideout,
		Title:      "NORMAN'S ROOFTOP HIDEOUT",
		Background: inc.BgRooftopHideout,
		Exits: []inkwell.Exit{
			{
				To:    inc.SceneNormanApartment,
				Label: "back down into the flat",
				Side:  inkwell.ExitNear,
			},
			{
				To:    inc.SceneBBSTerminal,
				Label: "the old terminal",
				Side:  inkwell.ExitRight,
			},
		},
	})
}
