package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	BackgroundManager.Register(Background{
		Name: BgSecretLab,
		Path: "assets/bg/secret_lab.png",
		Kind: inkwell.AssetImage,
	})
}
