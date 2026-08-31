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
	"realworld/inc/system/armilla"
	"realworld/inc/system/theme"
	"realworld/inc/system/widget"
	"realworld/inc/system/world"
)

type Options struct {
	Scene  string
	Finale bool
	Lang   string
	Debug  bool
}

var exitLook = inkwell.Text{
	{Lang: "en", Text: "That way: %s."},
	{Lang: "hu", Text: "Arra van: %s."},
}

func New(options Options) *inkwell.Game {
	start := options.Scene
	if start == "" {
		start = inc.SceneAlley
	}

	inkwell.Fallback = "en"
	inkwell.Lang = options.Lang

	scene.FillSelectorPins()

	game := inkwell.NewGame("Real World", widget.ScreenWidth, widget.ScreenHeight)
	game.MaxLogLines = 64
	game.WindowScale = 2
	game.SceneRect = inkwell.Rect(0, widget.TopBarHeight, widget.ScreenWidth, widget.HUDTop-widget.TopBarHeight)
	game.Player = inc.Paul
	game.Walkboxes = scene.Floor
	game.ExitLook = func(exit inkwell.Exit) inkwell.Action {
		return inkwell.Say(inc.Paul, inkwell.Format(exitLook, exit.Label))
	}
	game.ExitTake = func(inkwell.Exit) inkwell.Action {
		return inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "It's a way out, not a thing."},
			{Lang: "hu", Text: "Az egy kijárat, nem egy tárgy."},
		})
	}
	game.UseWithFail = script.UseWithFail

	game.AssetManager = background.BackgroundManager
	game.CharacterManager = character.CharacterManager
	game.DialogueManager = dialog.DialogManager
	game.ItemManager = item.ItemManager
	game.SceneManager = scene.SceneManager
	game.ScriptManager = script.ScriptManager
	game.WidgetManager = widget.WidgetManager
	theme.ThemeManager.Each(game.ThemeManager.Register)

	world.Attach(game)
	armilla.Reset()
	game.UseTheme(inc.RealWorld)
	if options.Debug || debugBuild {
		game.State.SetFlag(inc.FlagDebug)
	}

	game.StartAt(start)
	game.OnStart(inc.ScriptOpening)
	if options.Finale {
		game.OnFinale(inc.ScriptFinale)
	}
	return game
}
