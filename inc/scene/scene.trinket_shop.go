package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.SceneTrinketShop,
		Title:      "CHINATOWN — TRINKET SHOP",
		Background: inc.BgTrinketShop,
		Exits: []inkwell.Exit{
			exitToSelector,
			{
				To:    inc.SceneSmallRestaurant,
				Label: "the eating place next door",
				Side:  inkwell.ExitRight,
			},
			{
				To:      inc.SceneSecretClub,
				Label:   "the way in at the back",
				Side:    inkwell.ExitBack,
				Needs:   inc.FlagClubEntry,
				Blocked: inkwell.Say(inc.Paul, "Just a shop, as far as the man behind the counter is concerned."),
			},
		},
	})
}
