package scene

import ()

var Showroom = SceneManager.Register(&Scene{
	Name: "showroom",
	Title: Text{
		{Lang: "en", Text: "NEUMATRONIC SHOWROOM"},
		{Lang: "hu", Text: "NEUMATRONIC BEMUTATÓTEREM"},
	},
	Layers: []SceneLayer{
		{
			Asset: &Asset{
				Name: "showroom",
				Path: "assets/bg/showroom.png",
			},
			Z: -1,
		},
	},
})

var _ = Showroom.Leads(
	exitToSelector,
)
