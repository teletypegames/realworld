package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func backgroundHospital() inkwell.Asset {
	return inkwell.Asset{
		Name: BgHospital,
		Path: "assets/bg/hospital.png",
		Kind: inkwell.AssetImage,
	}
}
