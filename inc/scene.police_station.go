package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	SceneManager.Register(Scene{
		Name:       ScenePoliceStation,
		Title:      "SFPD — STATION AND HOLDING",
		Background: BgPoliceStation,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
