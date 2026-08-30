package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.SceneStreet,
		Title:      "STREET",
		Background: inc.BgStreet,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
