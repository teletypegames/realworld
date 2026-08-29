package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// columbarium — deck 23.
func columbarium() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgColumbarium,
		Path: "assets/bg/columbarium.png",
		Kind: inkwell.AssetImage,
	}
}
