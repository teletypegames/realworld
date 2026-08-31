package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name: inc.ScenePublicBBS,
		Title: inkwell.Text{
			{Lang: "en", Text: "PUBLIC BBS TERMINAL"},
			{Lang: "hu", Text: "NYILVÁNOS BBS TERMINÁL"},
		},
		Background: inc.BgPublicBBS,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
