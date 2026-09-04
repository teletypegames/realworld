package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

var Showroom = SceneManager.Register(&Scene{
	Name: "showroom",
	Title: inkwell.Text{
		{Lang: "en", Text: "NEUMATRONIC SHOWROOM"},
		{Lang: "hu", Text: "NEUMATRONIC BEMUTATÓTEREM"},
	},
})

var _ = Showroom.Leads(
	exitToSelector,
)
