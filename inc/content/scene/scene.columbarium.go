package scene

import ()

var Columbarium = SceneManager.Register(&Scene{
	Name: "columbarium",
	Title: Text{
		{Lang: "en", Text: "COLUMBARIUM — DEX'S MEMORIAL"},
		{Lang: "hu", Text: "KOLUMBÁRIUM — DEX EMLÉKHELYE"},
	},
	Layers: []SceneLayer{
		{
			Asset: &Asset{
				Name: "columbarium",
				Path: "assets/bg/columbarium.png",
			},
			Z: -1,
		},
	},
})

var _ = Columbarium.Leads(
	exitToSelector,
)
