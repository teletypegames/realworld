package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func backgroundPoliceStation() inkwell.Asset {
	return inkwell.Asset{
		Name: BgPoliceStation,
		Path: "assets/bg/police_station.png",
		Kind: inkwell.AssetImage,
	}
}
