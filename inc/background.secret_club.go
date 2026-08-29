package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func backgroundSecretClub() inkwell.Asset {
	return inkwell.Asset{
		Name: BgSecretClub,
		Path: "assets/bg/secret_club.png",
		Kind: inkwell.AssetImage,
	}
}
