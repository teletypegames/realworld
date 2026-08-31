package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name: inc.ScenePoliceStation,
		Title: inkwell.Text{
			{Lang: "en", Text: "SFPD — STATION AND HOLDING"},
			{Lang: "hu", Text: "SFPD — ŐRS ÉS FOGDA"},
		},
		Background: inc.BgPoliceStation,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
