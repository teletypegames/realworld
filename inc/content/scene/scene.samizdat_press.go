package scene

import ()

var SamizdatPress = SceneManager.Register(&Scene{
	Name: "samizdat_press",
	Title: Text{
		{Lang: "en", Text: "SAMIZDAT PRINTING HOUSE"},
		{Lang: "hu", Text: "SZAMIZDAT NYOMDA"},
	},
	Layers: []SceneLayer{
		{
			Asset: &Asset{
				Name: "samizdat_press",
				Path: "assets/bg/samizdat_press.png",
			},
			Z: -1,
		},
	},
})

var _ = SamizdatPress.Leads(
	exitToSelector,
)
