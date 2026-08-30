package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
	"git.teletypegames.org/games/realworld/inc/world"
)

func init() {
	Manager.Register(Script{
		Name: inc.ScriptTapeInsert,
		Actions: inkwell.Seq(
			world.Fn(func(ctx *inkwell.Ctx) {
				item := world.TakePending()
				if item == "" {
					return
				}
				world.SetSlot2(item)
				ctx.Game.Inventory.Remove(item)
				ctx.Game.State.SetVar(inc.VarArmillaStrip, "SLOT2: READING")
			}),
			inkwell.Say(inc.Paul, "Right. Let's see who you are."),
			inkwell.Say(inc.Dex, "Clicked in. Spinning. And now: nothing."),
			inkwell.Say(inc.Dex, "A stranger's tape in your own Armilla, Paul. I hope you know what you're doing."),
		),
	})
}
