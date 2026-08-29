// Package character registers the game's characters.
//
// Tapes are characters too, not props — the wiki is explicit about it. They
// have no body in the scene, so they carry no W/H and never appear in
// Scene.Actors; their voice lives in the tape channel (see world.TapeSay).
package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/world"
)

// Register adds every character to the game.
func Register(w *world.World) {
	for _, c := range []inkwell.Character{
		paul(),
		dex(),
		mysteryTape(),
		supportTape(),
	} {
		w.G.CharacterManager.Register(c)
	}
}
