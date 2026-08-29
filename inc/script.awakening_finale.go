package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func scriptAwakeningFinale(w *World) inkwell.Script {
	return inkwell.Script{
		Name: ScriptFinale,
		Actions: inkwell.Seq(
			SetMode(w, ModeCutscene),
			inkwell.Wait(0.6),
			inkwell.Say(Paul, "I'm putting it in. That's all this is."),
			inkwell.Wait(0.4),
			UseTheme(NokiaPunk),
			inkwell.Say("norman", "No — this isn't what you were supposed to do!"),
			inkwell.Wait(0.6),
			TapeSay(TapeMystery, "Thank you, Paul. You did exactly what I asked, the whole way through."),
			inkwell.Wait(1.0),
			inkwell.ShowEnd("REAL WORLD — end of game two"),
		),
	}
}
