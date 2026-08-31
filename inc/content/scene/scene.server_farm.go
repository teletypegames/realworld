package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	SceneManager.Register(Scene{
		Name: inc.SceneServerFarm,
		Title: inkwell.Text{
			{Lang: "en", Text: "SERVER FARM"},
			{Lang: "hu", Text: "SZERVERFARM"},
		},
		Background: inc.BgServerFarm,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
