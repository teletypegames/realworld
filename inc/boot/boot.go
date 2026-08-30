package boot

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
	"git.teletypegames.org/games/realworld/inc/background"
	"git.teletypegames.org/games/realworld/inc/character"
	"git.teletypegames.org/games/realworld/inc/dialog"
	"git.teletypegames.org/games/realworld/inc/item"
	"git.teletypegames.org/games/realworld/inc/scene"
	"git.teletypegames.org/games/realworld/inc/script"
	_ "git.teletypegames.org/games/realworld/inc/tape"
	"git.teletypegames.org/games/realworld/inc/theme"
	"git.teletypegames.org/games/realworld/inc/widget"
	"git.teletypegames.org/games/realworld/inc/world"
)

type Opts struct {
	Scene  string
	Finale bool
}

func New(o Opts) *inkwell.Game {
	start := o.Scene
	if start == "" {
		start = inc.SceneAlley
	}

	scene.FillSelectorPins()

	g := inkwell.NewGame("Real World", widget.ScreenW, widget.ScreenH)
	g.MaxLogLines = 64
	g.WindowScale = 2
	g.SceneRect = inkwell.Rect(0, widget.TopBarH, widget.ScreenW, widget.HUDTop-widget.TopBarH)
	g.Player = inc.Paul
	g.Walkboxes = scene.Floor
	g.ExitLook = func(e inkwell.Exit) inkwell.Action {
		return inkwell.Say(inc.Paul, "That way: "+e.Label+".")
	}
	g.ExitTake = func(inkwell.Exit) inkwell.Action {
		return inkwell.Say(inc.Paul, "It's a way out, not a thing.")
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
