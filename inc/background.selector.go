package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// selector — the wiki's Helyszínválasztó — no deck number, and not drawn.
//
// Not painted yet: the file is missing on purpose, and inkwell draws its
// placeholder instead — which is exactly the greybox convention.
func backgroundSelector() inkwell.Asset {
	return inkwell.Asset{
		Name: BgSelector,
		Path: "assets/bg/selector.png",
		Kind: inkwell.AssetImage,
	}
}
