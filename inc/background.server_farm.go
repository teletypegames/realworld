package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// serverFarm — deck 14.
func backgroundServerFarm() inkwell.Asset {
	return inkwell.Asset{
		Name: BgServerFarm,
		Path: "assets/bg/server_farm.png",
		Kind: inkwell.AssetImage,
	}
}
