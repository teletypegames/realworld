package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	BackgroundManager.Register(Background{
		Name: inc.BackgroundNormanApartment,
		Path: "assets/bg/norman_apartment.png",
		Kind: inkwell.AssetImage,
	})
}
