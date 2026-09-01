package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = SceneManager.Register(Scene{
	Name: inc.SceneNormanApartment,
	Title: inkwell.Text{
		{Lang: "en", Text: "NORMAN'S APARTMENT"},
		{Lang: "hu", Text: "NORMAN LAKÁSA"},
	},
	Exits: []inkwell.Exit{
		exitToSelector,
		{
			To: inc.SceneRooftopHideout,
			Label: inkwell.Text{
				{Lang: "en", Text: "the stairs up to the roof"},
				{Lang: "hu", Text: "a lépcső fel a tetőre"},
			},
			Side: inkwell.ExitBack,
		},
		{
			To: inc.SceneBBSTerminal,
			Label: inkwell.Text{
				{Lang: "en", Text: "Norman's terminal"},
				{Lang: "hu", Text: "Norman terminálja"},
			},
			Side: inkwell.ExitRight,
		},
	},
})
