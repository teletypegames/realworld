package inc

// Characters — the game's cast, one file each.
//
// Tapes are characters too, not props — the wiki is explicit about it. They
// have no body in the scene, so they carry no W/H and never appear in
// Scene.Actors; their voice lives in the tape channel (see TapeSay).

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// registerCharacter adds every character to the game.
func registerCharacter(w *World) {
	for _, c := range []inkwell.Character{
		characterPaul(),
		characterDex(),
		characterMysteryTape(),
		characterSupportTape(),
	} {
		w.G.CharacterManager.Register(c)
	}
}
