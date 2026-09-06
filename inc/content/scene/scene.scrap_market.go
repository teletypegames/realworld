package scene

import ()

var ScrapMarket = SceneManager.Register(&Scene{
	Name: "scrap_market",
	Title: Text{
		{Lang: "en", Text: "SCRAP MARKET"},
		{Lang: "hu", Text: "ÓCSKAPIAC"},
	},
	Layers: []SceneLayer{
		{
			Asset: &Asset{
				Name: "scrap_market",
				Path: "assets/bg/scrap_market.png",
			},
			Z: -1,
		},
	},
})

var _ = ScrapMarket.Leads(
	exitToSelector,
)
