package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgSecretClub,
		Path: "assets/bg/secret_club.png",
		Kind: inkwell.AssetImage,
	})
}
