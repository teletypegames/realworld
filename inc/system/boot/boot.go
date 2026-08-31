package boot

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/background"
	"realworld/inc/content/character"
	"realworld/inc/content/dialog"
	"realworld/inc/content/item"
	"realworld/inc/content/scene"
	"realworld/inc/content/script"
	_ "realworld/inc/content/tape"
	"realworld/inc/system/theme"
	"realworld/inc/system/widget"
	"realworld/inc/system/world"
)

type Opts struct {
	Scene  string
	Finale bool
	Lang   string
}

var exitLook = inkwell.Text{
	{Lang: "en", Text: "That way: %s."},
	{Lang: "hu", Text: "Arra van: %s."},
}

func New(o Opts) *inkwell.Game {
	start := o.Scene
	if start == "" {
		start = inc.SceneAlley
	}

	inkwell.Fallback = "en"
	inkwell.Lang = o.Lang

	scene.FillSelectorPins()

	g := inkwell.NewGame("Real World", widget.ScreenW, widget.ScreenH)
	g.MaxLogLines = 64
	g.WindowScale = 2
	g.SceneRect = inkwell.Rect(0, widget.TopBarH, widget.ScreenW, widget.HUDTop-widget.TopBarH)
	g.Player = inc.Paul
	g.Walkboxes = scene.Floor
	g.ExitLook = func(e inkwell.Exit) inkwell.Action {
		return inkwell.Say(inc.Paul, inkwell.Format(exitLook, e.Label))
	}
	g.ExitTake = func(inkwell.Exit) inkwell.Action {
		return inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "It's a way out, not a thing."},
			{Lang: "hu", Text: "Az egy kijárat, nem egy tárgy."},
		})
	}
	g.UseWithFail = script.UseWithFail

	g.AssetManager = background.Manager
	g.CharacterManager = character.Manager
	g.DialogueManager = dialog.Manager
	g.ItemManager = item.Manager
	g.SceneManager = scene.Manager
	g.ScriptManager = script.Manager
	g.WidgetManager = widget.Manager
	theme.Manager.Each(g.ThemeManager.Register)

	world.Attach(g)
	g.UseTheme(inc.RealWorld)

	g.StartAt(start)
	g.OnStart(inc.ScriptOpening)
	if o.Finale {
		g.OnFinale(inc.ScriptFinale)
	}
	return g
}
