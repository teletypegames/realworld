package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	SceneManager.Register(Scene{
		Name: inc.SceneSmallRestaurant,
		Title: inkwell.Text{
			{Lang: "en", Text: "SMALL RESTAURANT — NEXT DOOR"},
			{Lang: "hu", Text: "KIS ÉTTEREM — A SZOMSZÉDBAN"},
		},
		Background: inc.BgSmallRestaurant,
		Exits: []inkwell.Exit{
			{
				To: inc.SceneTrinketShop,
				Label: inkwell.Text{
					{Lang: "en", Text: "back to the trinket shop"},
					{Lang: "hu", Text: "vissza a csecsebecseboltba"},
				},
				Side: inkwell.ExitLeft,
			},
		},
	})
}
