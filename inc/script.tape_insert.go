package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func scriptTapeInsert(w *World) inkwell.Script {
	return inkwell.Script{
		Name: ScriptTapeInsert,
		Actions: inkwell.Seq(
			Fn(func(ctx *inkwell.Ctx) {
				item := w.TakePending()
				if item == "" {
					return
				}
				w.SetSlot2(item)
				ctx.Game.Inventory.Remove(item)
				ctx.Game.State.SetVar(VarArmillaStrip, "SLOT2: READING")
			}),
			inkwell.Say(Paul, "Right. Let's see who you are."),
			TapeSay(Dex, "Clicked in. Spinning. And now: nothing."),
			TapeSay(Dex, "A stranger's tape in your own Armilla, Paul. I hope you know what you're doing."),
		),
	}
}
