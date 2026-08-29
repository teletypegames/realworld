package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// policeStation — deck 04.
func policeStation() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgPoliceStation,
		Path: "assets/bg/police_station.png",
		Kind: inkwell.AssetImage,
	}
}
