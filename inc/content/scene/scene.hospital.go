package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = SceneManager.Register(Scene{
	Name: inc.SceneHospital,
	Title: inkwell.Text{
		{Lang: "en", Text: "HOSPITAL — PSYCHIATRIC WING"},
		{Lang: "hu", Text: "KÓRHÁZ — PSZICHIÁTRIAI SZÁRNY"},
	},
	Exits: []inkwell.Exit{
		exitToSelector,
	},
})
