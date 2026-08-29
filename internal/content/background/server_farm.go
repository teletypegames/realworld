package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// serverFarm — deck 14.
func serverFarm() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgServerFarm,
		Path: "assets/bg/server_farm.png",
		Kind: inkwell.AssetImage,
	}
}
