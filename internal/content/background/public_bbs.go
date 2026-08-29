package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// publicBbs — deck 17.
func publicBbs() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgPublicBBS,
		Path: "assets/bg/public_bbs.png",
		Kind: inkwell.AssetImage,
	}
}
