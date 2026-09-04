package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type bbsTerminal struct{}

var BBSTerminal = SceneManager.Register(&Scene{
	Name: "bbs_terminal",
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

func (bbsTerminal) Scene() *inkwell.Scene { return BBSTerminal }
