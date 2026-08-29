package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	BackgroundManager.Register(Background{
		Name: BgHospital,
		Path: "assets/bg/hospital.png",
		Kind: inkwell.AssetImage,
	})
}
