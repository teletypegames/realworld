package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Opts struct {
	Screen string
	Finale bool
}

func New(o Opts) *inkwell.Game {
	start := o.Screen
	if start == "" {
		start = ScreenAlley
	}

	g := inkwell.NewGame("Real World", ScreenW, ScreenH)
	g.MaxLogLines = 64

	w := newWorld(g)
	registerTheme(g)
	registerContent(w)

	g.UseTheme(RealWorld)
	registerUI(w)

	g.StartAt(start)
	g.OnStart(inkwell.Seq(bootOpening(w, start, o.Finale)...))
	return g
}

func bootOpening(w *World, start string, finale bool) []inkwell.Action {
	acts := []inkwell.Action{
		EnterScene(w, start),

		inkwell.SetFlag(FlagPoliceTip),
		inkwell.Say(Paul, "Back alley. Just like the clerk said."),
		TapeSay(Dex, "A government office tipped you off to remove something from a scene. That doesn't strike you as odd?"),
	}
	if finale {
		acts = append(acts, inkwell.RunScript(ScriptFinale))
	}
	return acts
}
