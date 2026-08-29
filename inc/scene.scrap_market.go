package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	SceneManager.Register(Scene{
		Name:       SceneScrapMarket,
		Title:      "SCRAP MARKET",
		Background: BgScrapMarket,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
