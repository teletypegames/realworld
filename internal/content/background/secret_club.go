package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// secretClub — deck 10.
func secretClub() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgSecretClub,
		Path: "assets/bg/secret_club.png",
		Kind: inkwell.AssetImage,
	}
}
