package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	BackgroundManager.Register(Background{
		Name: inc.BgSecretLab,
		Path: "assets/bg/secret_lab.png",
		Kind: inkwell.AssetImage,
	})
}
