package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	BackgroundManager.Register(Background{
		Name: BgNormanApartment,
		Path: "assets/bg/norman_apartment.png",
		Kind: inkwell.AssetImage,
	})
}
