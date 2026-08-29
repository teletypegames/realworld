package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func backgroundSecretLab() inkwell.Asset {
	return inkwell.Asset{
		Name: BgSecretLab,
		Path: "assets/bg/secret_lab.png",
		Kind: inkwell.AssetImage,
	}
}
