package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

var ScrapMarket = SceneManager.Register(&Scene{
	Name: "scrap_market",
	Title: inkwell.Text{
		{Lang: "en", Text: "SCRAP MARKET"},
		{Lang: "hu", Text: "ÓCSKAPIAC"},
	},
})

var _ = ScrapMarket.Leads(
	exitToSelector,
)
