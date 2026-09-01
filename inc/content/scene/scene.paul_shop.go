package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = SceneManager.Register(Scene{
	Name: inc.ScenePaulShop,
	Title: inkwell.Text{
		{Lang: "en", Text: "PAUL'S SHOP — JUNK AND GARAGE"},
		{Lang: "hu", Text: "PAUL BOLTJA — ÓCSKAVAS ÉS MŰHELY"},
	},
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
