package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// secretLab — deck 15.
func secretLab() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgSecretLab,
		Path: "assets/bg/secret_lab.png",
		Kind: inkwell.AssetImage,
	}
}
