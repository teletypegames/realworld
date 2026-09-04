package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type columbarium struct{}

var Columbarium = SceneManager.Register(&Scene{
	Name: "columbarium",
	Title: inkwell.Text{
		{Lang: "en", Text: "COLUMBARIUM — DEX'S MEMORIAL"},
		{Lang: "hu", Text: "KOLUMBÁRIUM — DEX EMLÉKHELYE"},
	},
	Exits: []inkwell.Exit{
		exitToSelector,
	},
})

func (columbarium) Scene() *inkwell.Scene { return Columbarium }
