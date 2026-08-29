package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	SceneManager.Register(Scene{
		Name:       ScenePublicBBS,
		Title:      "PUBLIC BBS TERMINAL",
		Background: BgPublicBBS,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
