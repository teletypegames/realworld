package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgHospital,
		Path: "assets/bg/hospital.png",
		Kind: inkwell.AssetImage,
	})
}
