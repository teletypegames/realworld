package scene

import ()

var Showroom = SceneManager.Register(&Scene{
	Name: "showroom",
	Title: Text{
		{Lang: "en", Text: "NEUMATRONIC SHOWROOM"},
		{Lang: "hu", Text: "NEUMATRONIC BEMUTATÓTEREM"},
	},
})

var _ = Showroom.Leads(
	exitToSelector,
)
