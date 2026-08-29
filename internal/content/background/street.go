package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// street — deck 12.
func street() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgStreet,
		Path: "assets/bg/street.png",
		Kind: inkwell.AssetImage,
	}
}
