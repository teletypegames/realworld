package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	SceneManager.Register(Scene{
		Name: inc.SceneBBSTerminal,
		Title: inkwell.Text{
			{Lang: "en", Text: "TERMINAL — BBS ACCESS POINT"},
			{Lang: "hu", Text: "TERMINÁL — BBS HOZZÁFÉRÉSI PONT"},
		},
		Exits: []inkwell.Exit{
			{
				Label: inkwell.Text{
					{Lang: "en", Text: "step back from the terminal"},
					{Lang: "hu", Text: "hátrébb lépni a terminálról"},
				},
				Side: inkwell.ExitNear,
			},
		},
	})
}
