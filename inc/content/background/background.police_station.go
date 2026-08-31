package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	BackgroundManager.Register(Background{
		Name: inc.BackgroundPoliceStation,
		Path: "assets/bg/police_station.png",
		Kind: inkwell.AssetImage,
	})
}
