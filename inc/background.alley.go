package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// alley — part of the catalogue's "Noodle háza & a sikátor" row — not in the deck.
//
// Not painted yet: the file is missing on purpose, and inkwell draws its
// placeholder instead — which is exactly the greybox convention.
func backgroundAlley() inkwell.Asset {
	return inkwell.Asset{
		Name: BgAlley,
		Path: "assets/bg/alley.png",
		Kind: inkwell.AssetImage,
	}
}
