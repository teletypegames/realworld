package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type scrapMarket struct{}

var ScrapMarket = SceneManager.Register(&Scene{
	Name: "scrap_market",
	Title: inkwell.Text{
		{Lang: "en", Text: "SCRAP MARKET"},
		{Lang: "hu", Text: "ÓCSKAPIAC"},
	},
	Exits: []inkwell.Exit{
		exitToSelector,
	},
})

func (scrapMarket) Scene() *inkwell.Scene { return ScrapMarket }
