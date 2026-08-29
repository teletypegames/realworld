package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	SceneManager.Register(Scene{
		Name:       SceneSmallRestaurant,
		Title:      "SMALL RESTAURANT — NEXT DOOR",
		Background: BgSmallRestaurant,
		Exits: []inkwell.Exit{
			{
				To:    SceneTrinketShop,
				Label: "back to the trinket shop",
				Side:  inkwell.ExitLeft,
			},
		},
	})
}
