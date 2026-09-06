package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

var BBSTerminal = SceneManager.Register(&Scene{
	Name: "bbs_terminal",
	Title: Text{
		{Lang: "en", Text: "TERMINAL — BBS ACCESS POINT"},
		{Lang: "hu", Text: "TERMINÁL — BBS HOZZÁFÉRÉSI PONT"},
	},
	Layers: []SceneLayer{
		{
			Asset: &Asset{
				Name: "bbs_terminal",
				Path: "assets/bg/bbs_terminal.png",
			},
			Z: -1,
		},
	},
})

var _ = BBSTerminal.Leads(
	Exit{
		Label: Text{
			{Lang: "en", Text: "step back from the terminal"},
			{Lang: "hu", Text: "hátrébb lépni a terminálról"},
		},
		Side: inkwell.ExitNear,
	},
)
