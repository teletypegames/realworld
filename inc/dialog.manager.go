package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func registerDialog(w *World) {
	for _, d := range []inkwell.Dialogue{
		dialogDexTalk(),
		dialogMysteryTapeSilent(),
	} {
		w.G.DialogueManager.Register(d)
	}
}
