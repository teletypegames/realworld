package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	BackgroundManager.Register(Background{
		Name: BgSecretClub,
		Path: "assets/bg/secret_club.png",
		Kind: inkwell.AssetImage,
	})
}
