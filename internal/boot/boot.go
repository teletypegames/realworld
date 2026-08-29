// Package boot wires the layers together: this is the only place where the
// world, the content, the theme and the HUD meet. main.go knows nothing else.
package boot

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/content"
	"git.teletypegames.org/games/realworld/internal/names"
	"git.teletypegames.org/games/realworld/internal/theme"
	"git.teletypegames.org/games/realworld/internal/ui"
	"git.teletypegames.org/games/realworld/internal/world"
)

// Opts are the run parameters.
type Opts struct {
	Scene  string // starting scene; empty means the alley
	Finale bool   // start with the finale, to try the Nokia-punk theme switch
}

// New builds a ready-to-run game.
func New(o Opts) *inkwell.Game {
	scene := o.Scene
	if scene == "" {
		scene = names.SceneAlley
	}

	g := inkwell.NewGame("Real World", ui.ScreenW, ui.ScreenH)
	g.MaxLogLines = 64

	w := world.New(g)
	theme.Register(g)
	content.Register(w)

	g.UseTheme(theme.RealWorld)
	ui.Register(w)

	g.StartAt(scene)
	g.OnStart(inkwell.Seq(opening(w, scene, o.Finale)...))
	return g
}

func opening(w *world.World, scene string, finale bool) []inkwell.Action {
	acts := []inkwell.Action{
		world.EnterScene(w, scene),
		// Beat 2 (the police station) will set the tip flag. Until that scene
		// exists we set it here so the alley slice is playable end to end —
		// the conditional branch itself is real code, not bypassed.
		inkwell.SetFlag(names.FlagPoliceTip),
		inkwell.Say(names.Paul, "Back alley. Just like the clerk said."),
		world.TapeSay(names.Dex, "A government office tipped you off to remove something from a scene. That doesn't strike you as odd?"),
	}
	if finale {
		acts = append(acts, inkwell.RunScript(names.ScriptFinale))
	}
	return acts
}
