package scene

import ()

var ScrapMarket = SceneManager.Register(&Scene{
	Name: "scrap_market",
	Title: Text{
		{Lang: "en", Text: "SCRAP MARKET"},
		{Lang: "hu", Text: "ÓCSKAPIAC"},
	},
})

var _ = ScrapMarket.Leads(
	exitToSelector,
)
