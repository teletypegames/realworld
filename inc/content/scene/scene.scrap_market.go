package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = SceneManager.Register(Scene{
	Name: inc.SceneScrapMarket,
	Title: inkwell.Text{
		{Lang: "en", Text: "SCRAP MARKET"},
		{Lang: "hu", Text: "ÓCSKAPIAC"},
	},
	Exits: []inkwell.Exit{
		exitToSelector,
	},
})
