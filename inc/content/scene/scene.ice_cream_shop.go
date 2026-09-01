package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = SceneManager.Register(Scene{
	Name: inc.SceneIceCreamShop,
	Title: inkwell.Text{
		{Lang: "en", Text: "ICE CREAM SHOP — WHERE THE BAR WAS"},
		{Lang: "hu", Text: "FAGYIZÓ — A RÉGI BÁR HELYÉN"},
	},
	Exits: []inkwell.Exit{
		exitToSelector,
	},
})
