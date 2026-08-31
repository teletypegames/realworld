package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	SceneManager.Register(Scene{
		Name: inc.SceneScrapMarket,
		Title: inkwell.Text{
			{Lang: "en", Text: "SCRAP MARKET"},
			{Lang: "hu", Text: "ÓCSKAPIAC"},
		},
		Background: inc.BgScrapMarket,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
