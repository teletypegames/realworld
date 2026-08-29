package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
	"git.teletypegames.org/games/realworld/internal/theme"
	"git.teletypegames.org/games/realworld/internal/world"
)

// awakeningFinale is the end of game two. The point of it, mechanically, is
// the theme switch: this is where the Nokia-punk texture breaks in.
//
// Run it early and often during development (`go run . -finale`) — it is the
// single most important visual beat in the game, so you want to be looking at
// it long before the surrounding scene exists.
func awakeningFinale(w *world.World) inkwell.Script {
	return inkwell.Script{
		Name: names.ScriptFinale,
		Actions: inkwell.Seq(
			world.SetMode(w, world.ModeCutscene),
			inkwell.Wait(0.6),
			inkwell.Say(names.Paul, "I'm putting it in. That's all this is."),
			inkwell.Wait(0.4),
			world.UseTheme(theme.NokiaPunk),
			inkwell.Say("norman", "No — this isn't what you were supposed to do!"),
			inkwell.Wait(0.6),
			world.TapeSay(names.TapeMystery, "Thank you, Paul. You did exactly what I asked, the whole way through."),
			inkwell.Wait(1.0),
			inkwell.ShowEnd("REAL WORLD — end of game two"),
		),
	}
}
