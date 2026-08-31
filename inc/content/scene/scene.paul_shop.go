package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name: inc.ScenePaulShop,
		Title: inkwell.Text{
			{Lang: "en", Text: "PAUL'S SHOP — JUNK AND GARAGE"},
			{Lang: "hu", Text: "PAUL BOLTJA — ÓCSKAVAS ÉS MŰHELY"},
		},
		Background: inc.BgPaulShop,
		Exits: []inkwell.Exit{
			{
				To: inc.SceneNoodleHouse,
				Label: inkwell.Text{
					{Lang: "en", Text: "the bus to San Francisco"},
					{Lang: "hu", Text: "a busz San Franciscóba"},
				},
				Side: inkwell.ExitNear,
			},
		},
	})
}
