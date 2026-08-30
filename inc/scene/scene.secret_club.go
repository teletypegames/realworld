package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.SceneSecretClub,
		Title:      "SECRET CLUB — THE UNDERWATER SUN",
		Background: inc.BgSecretClub,
		Exits: []inkwell.Exit{
			{
				To:    inc.SceneBBSTerminal,
				Label: "the terminal in the corner",
				Side:  inkwell.ExitBack,
			},
			{
				To:    inc.SceneTrinketShop,
				Label: "back out through the shop",
				Side:  inkwell.ExitNear,
			},
		},
	})
}
