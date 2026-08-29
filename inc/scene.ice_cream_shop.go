package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	SceneManager.Register(Scene{
		Name:       SceneIceCreamShop,
		Title:      "ICE CREAM SHOP — WHERE THE BAR WAS",
		Background: BgIceCreamShop,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
