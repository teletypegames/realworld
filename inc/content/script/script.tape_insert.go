package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/armilla"
	"realworld/inc/system/world"
)

func init() {
	ScriptManager.Register(Script{
		Name: inc.ScriptTapeInsert,
		Actions: inkwell.Seq(
			world.Func(func(context *inkwell.Ctx) {
				item := armilla.TakePending()
				if item == "" {
					return
				}
				armilla.SetSlot2(item)
				context.Game.Inventory.Remove(item)
				armilla.SetStrip("SLOT2: READING")
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
