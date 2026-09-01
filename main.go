package main

import (
	"embed"
	"flag"
	"log"

	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/background"
	"realworld/inc/content/character"
	"realworld/inc/content/dialog"
	"realworld/inc/content/item"
	"realworld/inc/content/scene"
	"realworld/inc/content/script"
	_ "realworld/inc/content/tape"
	"realworld/inc/system"
	"realworld/inc/system/theme"
	"realworld/inc/system/widget"
)

//go:embed assets
var assets embed.FS

func init() { inkwell.AssetFS = assets }

func main() {
	start := flag.String("scene", "", "starting scene (empty: the alley)")
	finale := flag.Bool("finale", false, "start with the finale, to try the Nokia-punk theme switch")
	lang := flag.String("lang", "en", "language the game speaks (en, hu)")
	debug := flag.Bool("debug", false, "open the debug menu with d (a debug build has it anyway)")
	flag.Parse()

	if *start == "" {
		*start = inc.SceneAlley
	}

	inkwell.Fallback = "en"
	inkwell.Lang = *lang

	game := inkwell.NewGame("Real World", widget.ScreenWidth, widget.ScreenHeight)
	game.MaxLogLines = 64
	game.WindowScale = 2
	game.SceneRect = inkwell.Rect(0, widget.TopBarHeight, widget.ScreenWidth, widget.HUDTop-widget.TopBarHeight)
	game.Player = inc.Paul
	game.Walkboxes = scene.Floor
	game.ExitLook = script.ExitLook
	game.ExitTake = script.ExitTake
	game.UseWithFail = script.UseWithFail

	game.AssetManager = background.BackgroundManager
	game.CharacterManager = character.CharacterManager
	game.DialogManager = dialog.DialogManager
	game.ItemManager = item.ItemManager
	game.SceneManager = scene.SceneManager
	game.ScriptManager = script.ScriptManager
	game.WidgetManager = widget.WidgetManager
	theme.ThemeManager.Each(game.ThemeManager.Register)

	system.System.Attach(game)
	game.UseTheme(inc.RealWorld)
	if *debug || debugBuild {
		game.State.SetFlag(inc.FlagDebug)
	}

	game.StartAt(*start)
	game.OnStart(inc.ScriptOpening)
	if *finale {
		game.OnFinale(inc.ScriptFinale)
	}

	if err := inkwell.Run(game); err != nil {
		log.Fatal(err)
	}
}
