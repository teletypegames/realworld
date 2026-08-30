package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Background{
		Name: inc.BgSecretLab,
		Path: "assets/bg/secret_lab.png",
		Kind: inkwell.AssetImage,
	})
}
