package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	SceneManager.Register(Scene{
		Name:       SceneTrinketShop,
		Title:      "CHINATOWN — TRINKET SHOP",
		Background: BgTrinketShop,
		Exits: []inkwell.Exit{
			exitToSelector,
			{
				To:    SceneSmallRestaurant,
				Label: "the eating place next door",
				Side:  inkwell.ExitRight,
			},
			{
				To:      SceneSecretClub,
				Label:   "the way in at the back",
				Side:    inkwell.ExitBack,
				Needs:   FlagClubEntry,
				Blocked: inkwell.Say(Paul, "Just a shop, as far as the man behind the counter is concerned."),
			},
		},
	})
}
