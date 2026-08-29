package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
	"git.teletypegames.org/games/realworld/internal/world"
)

// tapeInsert runs whenever a tape goes into slot 2: it activates the tape and
// makes Dex comment. The wiki calls the comment mandatory — Dex remarks on
// every newly loaded tape, without exception.
func tapeInsert(w *world.World) inkwell.Script {
	return inkwell.Script{
		Name: names.ScriptTapeInsert,
		Actions: inkwell.Seq(
			world.Fn(func(ctx *inkwell.Ctx) {
				item := w.TakePending()
				if item == "" {
					return
				}
				w.SetSlot2(item)
				ctx.Game.Inventory.Remove(item)
				ctx.Game.State.SetVar(names.VarArmillaStrip, "SLOT2: READING")
			}),
			inkwell.Say(names.Paul, "Right. Let's see who you are."),
			world.TapeSay(names.Dex, "Clicked in. Spinning. And now: nothing."),
			world.TapeSay(names.Dex, "A stranger's tape in your own Armilla, Paul. I hope you know what you're doing."),
		),
	}
}
