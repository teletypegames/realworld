package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type showroom struct{}

var Showroom = SceneManager.Register(&Scene{
	Name: "showroom",
	Title: inkwell.Text{
		{Lang: "en", Text: "NEUMATRONIC SHOWROOM"},
		{Lang: "hu", Text: "NEUMATRONIC BEMUTATÓTEREM"},
	},
	Exits: []inkwell.Exit{
		exitToSelector,
	},
})

func (showroom) Scene() *inkwell.Scene { return Showroom }
