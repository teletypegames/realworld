package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgServerFarm,
		Path: "assets/bg/server_farm.png",
		Kind: inkwell.AssetImage,
	})
}
