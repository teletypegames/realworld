package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// rooftopHideout — deck 16.
func rooftopHideout() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgRooftopHideout,
		Path: "assets/bg/rooftop_hideout.png",
		Kind: inkwell.AssetImage,
	}
}
