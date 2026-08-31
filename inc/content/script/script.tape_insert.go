package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/world"
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
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "Right. Let's see who you are."},
				{Lang: "hu", Text: "Jól van. Lássuk, ki vagy."},
			}),
			inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "Clicked in. Spinning. And now: nothing."},
				{Lang: "hu", Text: "Bekattant. Forog. És most: semmi."},
			}),
			inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "A stranger's tape in your own Armilla, Paul. I hope you know what you're doing."},
				{Lang: "hu", Text: "Egy idegen kazettája a saját Armilládban, Paul. Remélem, tudod, mit csinálsz."},
			}),
		),
	})
}
