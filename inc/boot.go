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
	TapeManager       = inkwell.NewManager[Tape]()
	ThemeManager      = inkwell.NewManager[Theme]()
	WidgetManager     = inkwell.NewManager[Widget]()
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

	fillSelectorPins()

	g := inkwell.NewGame("Real World", ScreenW, ScreenH)
	g.MaxLogLines = 64
	g.WindowScale = 2
	g.SceneRect = inkwell.Rect(0, TopBarH, ScreenW, HUDTop-TopBarH)
	g.Player = Paul
	g.Walkboxes = sceneFloor
	g.ExitLook = func(e inkwell.Exit) inkwell.Action {
		return inkwell.Say(Paul, "That way: "+e.Label+".")
	}
	g.ExitTake = func(inkwell.Exit) inkwell.Action {
		return inkwell.Say(Paul, "It's a way out, not a thing.")
	}
	g.UseWithFail = useWithFail

	g.AssetManager = BackgroundManager
	g.CharacterManager = CharacterManager
	g.DialogueManager = DialogManager
	g.ItemManager = ItemManager
	g.SceneManager = SceneManager
	g.ScriptManager = ScriptManager
	g.WidgetManager = WidgetManager
	ThemeManager.Each(g.ThemeManager.Register)

	World.attach(g)
	g.UseTheme(RealWorld)

	g.StartAt(start)
	g.OnStart(ScriptOpening)
	if o.Finale {
		g.OnFinale(ScriptFinale)
	}
	return g
}
