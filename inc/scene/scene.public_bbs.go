package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.ScenePublicBBS,
		Title:      "PUBLIC BBS TERMINAL",
		Background: inc.BgPublicBBS,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
