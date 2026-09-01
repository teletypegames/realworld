package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = SceneManager.Register(Scene{
	Name: inc.ScenePublicBBS,
	Title: inkwell.Text{
		{Lang: "en", Text: "PUBLIC BBS TERMINAL"},
		{Lang: "hu", Text: "NYILVÁNOS BBS TERMINÁL"},
	},
	Exits: []inkwell.Exit{
		exitToSelector,
	},
})
