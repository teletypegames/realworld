package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

var PublicBBS = SceneManager.Register(&Scene{
	Name: "public_bbs",
	Title: inkwell.Text{
		{Lang: "en", Text: "PUBLIC BBS TERMINAL"},
		{Lang: "hu", Text: "NYILVÁNOS BBS TERMINÁL"},
	},
})

var _ = PublicBBS.Leads(
	exitToSelector,
)
