package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.SceneHospital,
		Title:      "HOSPITAL — PSYCHIATRIC WING",
		Background: inc.BgHospital,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
