package scene

import ()

var Columbarium = SceneManager.Register(&Scene{
	Name: "columbarium",
	Title: Text{
		{Lang: "en", Text: "COLUMBARIUM — DEX'S MEMORIAL"},
		{Lang: "hu", Text: "KOLUMBÁRIUM — DEX EMLÉKHELYE"},
	},
})

var _ = Columbarium.Leads(
	exitToSelector,
)
