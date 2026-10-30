package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/script"
)

var Opening = SceneManager.Register(&Scene{
	Name: "opening",
	Title: Text{
		{Lang: "en", Text: "OPENING — PAUL RECIVES A LETTER FROM NOODLE"},
		{Lang: "hu", Text: "MEGNYITÓ — PAUL LEVELET KAP NOODLE-TŐL"},
	},
	Layers: []SceneLayer{
		{
			Asset: &Asset{
				Name: "opening",
				Path: "assets/bg/intro.png",
			},
			Z: -1,
		},
	},
	Triggers: []Trigger{
		{
			Name: "opening",
			When: Given{
				NotFlag: inc.FlagOpeningDone,
			},
			Do:   inkwell.RunScript(script.Opening),
			Once: true,
		},
	},
})

var _ = Opening.Leads(
	Exit{
		To: Bus,
		Label: Text{
			{Lang: "en", Text: "get on the bus"},
			{Lang: "hu", Text: "felszállni a buszra"},
		},
		Side: inkwell.ExitRight,
	},
)
