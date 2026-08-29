package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	SceneManager.Register(Scene{
		Name:       SceneSecretClub,
		Title:      "SECRET CLUB — THE UNDERWATER SUN",
		Background: BgSecretClub,
		Exits: []inkwell.Exit{
			{
				To:    SceneBBSTerminal,
				Label: "the terminal in the corner",
				Side:  inkwell.ExitBack,
			},
			{
				To:    SceneTrinketShop,
				Label: "back out through the shop",
				Side:  inkwell.ExitNear,
			},
		},
	})
}
