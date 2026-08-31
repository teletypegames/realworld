package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	SceneManager.Register(Scene{
		Name: inc.SceneSecretClub,
		Title: inkwell.Text{
			{Lang: "en", Text: "SECRET CLUB — THE UNDERWATER SUN"},
			{Lang: "hu", Text: "TITKOS KLUB — A VÍZ ALATTI NAP"},
		},
		Background: inc.BackgroundSecretClub,
		Exits: []inkwell.Exit{
			{
				To: inc.SceneBBSTerminal,
				Label: inkwell.Text{
					{Lang: "en", Text: "the terminal in the corner"},
					{Lang: "hu", Text: "a terminál a sarokban"},
				},
				Side: inkwell.ExitBack,
			},
			{
				To: inc.SceneTrinketShop,
				Label: inkwell.Text{
					{Lang: "en", Text: "back out through the shop"},
					{Lang: "hu", Text: "vissza a bolton át"},
				},
				Side: inkwell.ExitNear,
			},
		},
	})
}
