package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
	"git.teletypegames.org/games/realworld/inc/world"
)

func init() {
	Manager.Register(Script{
		Name: inc.ScriptFinale,
		Actions: inkwell.Seq(
			world.SetMode(world.ModeCutscene),
			inkwell.Wait(0.6),
			inkwell.Say(inc.Paul, "I'm putting it in. That's all this is."),
			inkwell.Wait(0.4),
			world.UseTheme(inc.NokiaPunk),
			inkwell.Say("norman", "No — this isn't what you were supposed to do!"),
			inkwell.Wait(0.6),
			inkwell.Say(inc.TapeMystery, "Thank you, Paul. You did exactly what I asked, the whole way through."),
			inkwell.Wait(1.0),
			inkwell.ShowEnd("REAL WORLD — end of game two"),
		),
	})
}
