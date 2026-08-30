package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.SceneBBSTerminal,
		Title:      "TERMINAL — BBS ACCESS POINT",
		Background: inc.BgBBSTerminal,
		Exits: []inkwell.Exit{
			{
				Label: "step back from the terminal",
				Side:  inkwell.ExitNear,
			},
		},
	})
}
