package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	SceneManager.Register(Scene{
		Name: inc.SceneNoodleHouse,
		Title: inkwell.Text{
			{Lang: "en", Text: "NOODLE'S HOUSE — SEALED"},
			{Lang: "hu", Text: "NOODLE HÁZA — LEPECSÉTELVE"},
		},
		Background: inc.BackgroundNoodleHouse,
		Exits: []inkwell.Exit{
			exitToSelector,
			{
				To: inc.SceneAlley,
				Label: inkwell.Text{
					{Lang: "en", Text: "the alley behind the house"},
					{Lang: "hu", Text: "a sikátor a ház mögött"},
				},
				Side: inkwell.ExitRight,
			},
		},
	})
}
