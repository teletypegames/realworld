// Package asset registers the game's assets.
//
// During the greybox phase no asset file exists on disk: inkwell draws a
// placeholder when a sprite is missing, which is exactly the wireframe deck's
// dashed-outline convention, for free. Only the registration has to exist —
// Game.Validate requires every scene background to name a registered asset.
package asset

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/internal/world"
)

// Register adds every asset to the game.
func Register(w *world.World) {
	for _, a := range []inkwell.Asset{
		alleyBackground(),
	} {
		w.G.AssetManager.Register(a)
	}
}
