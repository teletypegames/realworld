package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

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
