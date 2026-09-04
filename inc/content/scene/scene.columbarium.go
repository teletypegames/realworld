package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

var Columbarium = SceneManager.Register(&Scene{
	Name: "columbarium",
	Title: inkwell.Text{
		{Lang: "en", Text: "COLUMBARIUM — DEX'S MEMORIAL"},
		{Lang: "hu", Text: "KOLUMBÁRIUM — DEX EMLÉKHELYE"},
	},
})

var _ = Columbarium.Leads(
	exitToSelector,
)
