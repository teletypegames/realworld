package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.SceneSmallRestaurant,
		Title:      "SMALL RESTAURANT — NEXT DOOR",
		Background: inc.BgSmallRestaurant,
		Exits: []inkwell.Exit{
			{
				To:    inc.SceneTrinketShop,
				Label: "back to the trinket shop",
				Side:  inkwell.ExitLeft,
			},
		},
	})
}
