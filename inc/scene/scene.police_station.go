package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.ScenePoliceStation,
		Title:      "SFPD — STATION AND HOLDING",
		Background: inc.BgPoliceStation,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
