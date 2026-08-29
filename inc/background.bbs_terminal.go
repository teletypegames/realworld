package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// bbsTerminal — deck 11.
func backgroundBbsTerminal() inkwell.Asset {
	return inkwell.Asset{
		Name: BgBBSTerminal,
		Path: "assets/bg/bbs_terminal.png",
		Kind: inkwell.AssetImage,
	}
}
