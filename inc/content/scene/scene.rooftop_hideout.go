package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	SceneManager.Register(Scene{
		Name: inc.SceneRooftopHideout,
		Title: inkwell.Text{
			{Lang: "en", Text: "NORMAN'S ROOFTOP HIDEOUT"},
			{Lang: "hu", Text: "NORMAN BÚVÓHELYE A TETŐN"},
		},
		Exits: []inkwell.Exit{
			{
				To: inc.SceneNormanApartment,
				Label: inkwell.Text{
					{Lang: "en", Text: "back down into the flat"},
					{Lang: "hu", Text: "vissza le a lakásba"},
				},
				Side: inkwell.ExitNear,
			},
			{
				To: inc.SceneBBSTerminal,
				Label: inkwell.Text{
					{Lang: "en", Text: "the old terminal"},
					{Lang: "hu", Text: "a régi terminál"},
				},
				Side: inkwell.ExitRight,
			},
		},
	})
}
