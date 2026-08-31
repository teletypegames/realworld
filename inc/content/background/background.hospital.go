package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	BackgroundManager.Register(Background{
		Name: inc.BackgroundHospital,
		Path: "assets/bg/hospital.png",
		Kind: inkwell.AssetImage,
	})
}
