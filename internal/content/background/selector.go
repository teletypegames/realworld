package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// selector — the wiki's Helyszínválasztó — no deck number, and not drawn.
//
// Not painted yet: the file is missing on purpose, and inkwell draws its
// placeholder instead — which is exactly the greybox convention.
func selector() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgSelector,
		Path: "assets/bg/selector.png",
		Kind: inkwell.AssetImage,
	}
}
