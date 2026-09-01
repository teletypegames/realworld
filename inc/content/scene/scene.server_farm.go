package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = SceneManager.Register(Scene{
	Name: inc.SceneServerFarm,
	Title: inkwell.Text{
		{Lang: "en", Text: "SERVER FARM"},
		{Lang: "hu", Text: "SZERVERFARM"},
	},
	Exits: []inkwell.Exit{
		exitToSelector,
	},
})
