package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func registerScript(w *World) {
	for _, s := range []inkwell.Script{
		scriptTapeInsert(w),
		scriptAwakeningFinale(w),
	} {
		w.G.ScriptManager.Register(s)
	}
}
