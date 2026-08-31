package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
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
