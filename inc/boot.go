package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

var (
	BackgroundManager = inkwell.NewManager[Background]()
	CharacterManager  = inkwell.NewManager[Character]()
	DialogManager     = inkwell.NewManager[Dialog]()
	ItemManager       = inkwell.NewManager[Item]()
	SceneManager      = inkwell.NewManager[Scene]()
	ScriptManager     = inkwell.NewManager[Script]()
	ThemeManager      = inkwell.NewManager[Theme]()
)

type Opts struct {
	Scene  string
	Finale bool
}

func New(o Opts) *inkwell.Game {
	start := o.Scene
	if start == "" {
		start = SceneAlley
	}

	prepareScene()

	g := inkwell.NewGame("Real World", ScreenW, ScreenH)
	g.MaxLogLines = 64
	g.SceneRect = inkwell.Rect(0, TopBarH, ScreenW, HUDTop-TopBarH)
	g.ExitLook = func(e inkwell.Exit) inkwell.Action {
		return inkwell.Say(Paul, "That way: "+e.Label+".")
	}
	g.ExitTake = func(inkwell.Exit) inkwell.Action {
		return inkwell.Say(Paul, "It's a way out, not a thing.")
	}

	g.AssetManager = BackgroundManager
	g.CharacterManager = CharacterManager
	g.DialogueManager = DialogManager
	g.ItemManager = ItemManager
	g.SceneManager = SceneManager
	g.ScriptManager = ScriptManager
	ThemeManager.Each(g.ThemeManager.Register)

	World.attach(g)
	g.UseTheme(RealWorld)
	registerUI()

	g.StartAt(start)
	g.OnStart(inkwell.Seq(bootOpening(start, o.Finale)...))
	return g
}

func bootOpening(start string, finale bool) []inkwell.Action {
	acts := []inkwell.Action{
		inkwell.SetFlag(FlagPoliceTip),
		inkwell.Say(Paul, "Back alley. Just like the clerk said."),
		TapeSay(Dex, "A government office tipped you off to remove something from a scene. That doesn't strike you as odd?"),
	}
	if finale {
		acts = append(acts, inkwell.RunScript(ScriptFinale))
	}
	return acts
}
