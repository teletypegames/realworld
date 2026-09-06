package scene

import ()

var PublicBBS = SceneManager.Register(&Scene{
	Name: "public_bbs",
	Title: Text{
		{Lang: "en", Text: "PUBLIC BBS TERMINAL"},
		{Lang: "hu", Text: "NYILVÁNOS BBS TERMINÁL"},
	},
	Layers: []SceneLayer{
		{
			Asset: &Asset{
				Name: "public_bbs",
				Path: "assets/bg/public_bbs.png",
			},
			Z: -1,
		},
	},
})

var _ = PublicBBS.Leads(
	exitToSelector,
)
