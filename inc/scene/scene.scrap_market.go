package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.SceneScrapMarket,
		Title:      "SCRAP MARKET",
		Background: inc.BgScrapMarket,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
