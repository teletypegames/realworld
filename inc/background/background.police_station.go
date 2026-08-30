package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgPoliceStation,
		Path: "assets/bg/police_station.png",
		Kind: inkwell.AssetImage,
	})
}
