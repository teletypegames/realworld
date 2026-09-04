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
	"realworld/inc/content/speaker"
	"realworld/inc/content/stage"
	_ "realworld/inc/content/tape"
	"realworld/inc/system"
	"realworld/inc/system/theme"
	"realworld/inc/system/widget"
)

//go:embed assets
var assets embed.FS

func init() { inkwell.AssetFS = assets }

func main() {
	chosen := parseOptions()

	inkwell.Fallback = "en"
	inkwell.Lang = chosen.language

	game := chosen.game()
	system.System.Attach(game)

	if err := inkwell.Run(game); err != nil {
		log.Fatal(err)
	}
}

type options struct {
	scene    string
	finale   bool
	language string
	debug    bool
}

func parseOptions() options {
	startScene := flag.String("scene", "", "start straight in this scene, skipping the main menu")
	finale := flag.Bool("finale", false, "start with the finale, to try the Nokia-punk theme switch")
	language := flag.String("lang", "", "language the game speaks for this run (en, hu); the saved setting otherwise")
	debug := flag.Bool("debug", false, "open the debug menu with d (a debug build has it anyway)")
	flag.Parse()

	return options{
		scene:    *startScene,
		finale:   *finale,
		language: *language,
		debug:    *debug,
	}
}

func (chosen options) startStage() *stage.Stage {
	if chosen.scene != "" || chosen.finale {
		return stage.Game
	}
	return stage.SplashStudio
}

func (chosen options) startScene() *scene.Scene {
	if chosen.scene == "" {
		return scene.PaulShop
	}
	return scene.SceneManager.MustGet(chosen.scene)
}

func (chosen options) finaleScript() *script.Script {
	if !chosen.finale {
		return nil
	}
	return script.Finale
}

func (chosen options) flags() []string {
	if !chosen.debug && !debugBuild {
		return nil
	}
	return []string{inc.FlagDebug}
}

func (chosen options) game() *inkwell.Game {
	return inkwell.New(inkwell.Game{
		Title:       "Real World",
		App:         "realworld",
		Width:       widget.ScreenWidth,
		Height:      widget.ScreenHeight,
		WindowScale: 2,
		MaxLogLines: 64,
		SceneRect:   inkwell.Rect(0, widget.TopBarHeight, widget.ScreenWidth, widget.HUDTop-widget.TopBarHeight),

		StartStage:   chosen.startStage(),
		StartScene:   chosen.startScene(),
		StartScript:  script.Opening,
		FinaleScript: chosen.finaleScript(),
		StartTheme:   theme.RealWorld,
		Flags:        chosen.flags(),

		Player:      character.Paul,
		WalkOnClick: true,
		Walkboxes:   scene.Floor,
		ExitLook:    script.ExitLook,
		ExitTake:    script.ExitTake,
		UseWithFail: script.UseWithFail,

		AssetManager:     background.BackgroundManager,
		CharacterManager: character.CharacterManager,
		DialogManager:    dialog.DialogManager,
		ItemManager:      item.ItemManager,
		SceneManager:     scene.SceneManager,
		ScriptManager:    script.ScriptManager,
		SpeakerManager:   speaker.SpeakerManager,
		StageManager:     stage.StageManager,
		ThemeManager:     theme.ThemeManager,
		WidgetManager:    widget.WidgetManager,
	})
}
