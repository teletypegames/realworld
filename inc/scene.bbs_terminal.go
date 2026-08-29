package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	SceneManager.Register(Scene{
		Name:       SceneBBSTerminal,
		Title:      "TERMINAL — BBS ACCESS POINT",
		Background: BgBBSTerminal,
		Exits: []inkwell.Exit{
			{
				Label: "step back from the terminal",
				Side:  inkwell.ExitNear,
			},
		},
	})
}
