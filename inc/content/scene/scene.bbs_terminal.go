package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

var BBSTerminal = SceneManager.Register(&Scene{
	Name: "bbs_terminal",
	Title: inkwell.Text{
		{Lang: "en", Text: "TERMINAL — BBS ACCESS POINT"},
		{Lang: "hu", Text: "TERMINÁL — BBS HOZZÁFÉRÉSI PONT"},
	},
})

var _ = BBSTerminal.Leads(
	inkwell.Exit{
		Label: inkwell.Text{
			{Lang: "en", Text: "step back from the terminal"},
			{Lang: "hu", Text: "hátrébb lépni a terminálról"},
		},
		Side: inkwell.ExitNear,
	},
)
