package scene

import ()

var CuratorShop = SceneManager.Register(&Scene{
	Name: "curator_shop",
	Title: Text{
		{Lang: "en", Text: "CURATOR SERVICE — THE MASTER'S WORKSHOP"},
		{Lang: "hu", Text: "KURÁTORSZOLGÁLAT — A MESTER MŰHELYE"},
	},
	Layers: []SceneLayer{
		{
			Asset: &Asset{
				Name: "curator_shop",
				Path: "assets/bg/curator_shop.png",
			},
			Z: -1,
		},
	},
})

var _ = CuratorShop.Leads(
	exitToSelector,
)
