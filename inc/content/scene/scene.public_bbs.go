package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type publicBBS struct{}

var PublicBBS = SceneManager.Register(&Scene{
	Name: "public_bbs",
	Title: inkwell.Text{
		{Lang: "en", Text: "PUBLIC BBS TERMINAL"},
		{Lang: "hu", Text: "NYILVÁNOS BBS TERMINÁL"},
	},
	Exits: []inkwell.Exit{
		exitToSelector,
	},
})

func (publicBBS) Scene() *inkwell.Scene { return PublicBBS }
