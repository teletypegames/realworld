package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	BackgroundManager.Register(Background{
		Name: BgPoliceStation,
		Path: "assets/bg/police_station.png",
		Kind: inkwell.AssetImage,
	})
}
