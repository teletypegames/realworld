package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.SceneIceCreamShop,
		Title:      "ICE CREAM SHOP — WHERE THE BAR WAS",
		Background: inc.BgIceCreamShop,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
