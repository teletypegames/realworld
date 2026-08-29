package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// hackerspace — deck 06.
func hackerspace() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgHackerspace,
		Path: "assets/bg/hackerspace.png",
		Kind: inkwell.AssetImage,
	}
}
