package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// bbsTerminal — deck 11.
func bbsTerminal() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgBBSTerminal,
		Path: "assets/bg/bbs_terminal.png",
		Kind: inkwell.AssetImage,
	}
}
