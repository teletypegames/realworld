package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// samizdatPress — deck 21.
func samizdatPress() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgSamizdatPress,
		Path: "assets/bg/samizdat_press.png",
		Kind: inkwell.AssetImage,
	}
}
