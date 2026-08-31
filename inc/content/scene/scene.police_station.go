package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	SceneManager.Register(Scene{
		Name: inc.ScenePoliceStation,
		Title: inkwell.Text{
			{Lang: "en", Text: "SFPD — STATION AND HOLDING"},
			{Lang: "hu", Text: "SFPD — ŐRS ÉS FOGDA"},
		},
		Background: inc.BackgroundPoliceStation,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
