package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// alley — part of the catalogue's "Noodle háza & a sikátor" row — not in the deck.
//
// Not painted yet: the file is missing on purpose, and inkwell draws its
// placeholder instead — which is exactly the greybox convention.
func alley() inkwell.Asset {
	return inkwell.Asset{
		Name: names.BgAlley,
		Path: "assets/bg/alley.png",
		Kind: inkwell.AssetImage,
	}
}
