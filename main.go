package main

import (
	"embed"
	"flag"
	"log"

	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/achievement"
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

	if err := chosen.run(game); err != nil {
		log.Fatal(err)
	}
}

type options struct {
	scene    string
	finale   bool
	language string
	debug    bool
	explore  bool
}

func parseOptions() options {
	startScene := flag.String("scene", "", "start straight in this scene, skipping the main menu")
	finale := flag.Bool("finale", false, "start with the finale, to try the Nokia-punk theme switch")
	language := flag.String("lang", "", "language the game speaks for this run (en, hu); the saved setting otherwise")
	debug := flag.Bool("debug", false, "open the debug menu with d (a debug build has it anyway)")
	explore := flag.Bool("explore", false, "browse the game's entities instead of playing it")
	flag.Parse()

	return options{
		scene:    *startScene,
		finale:   *finale,
		language: *language,
		debug:    *debug,
		explore:  *explore,
	}
}

func (chosen options) run(game *inkwell.Game) error {
	if chosen.explore {
		return inkwell.Explore(game)
	}
	return inkwell.Run(game)
}

func (chosen options) startStage() *stage.Stage {
	if chosen.scene != "" || chosen.finale {
		return stage.Game
	}
	return stage.SplashStudio
}

func (chosen options) startScene() *scene.Scene {
	if chosen.scene == "" {
		return scene.PaulDoor
	}
	return scene.SceneManager.MustGet(chosen.scene)
}

func (chosen options) flags() []string {
	if !chosen.debug && !debugBuild {
		return nil
	}
	return []string{inc.FlagDebug}
}

func (chosen options) game() *inkwell.Game {
	return inkwell.New(inkwell.Game{
		Title:       "Underwater Sun Rising",
		App:         "realworld",
		Width:       widget.ScreenWidth,
		Height:      widget.ScreenHeight,
		WindowScale: 2,
		MaxLogLines: 64,
		SceneRect:   inkwell.Rect(0, widget.TopBarHeight, widget.ScreenWidth, widget.HUDTop-widget.TopBarHeight),

		StartStage: chosen.startStage(),
		StartScene: chosen.startScene(),
		StartTheme: theme.RealWorld,
		Flags:      chosen.flags(),
		Extensions: platformExtensions,

		Player:      character.Paul,
		WalkOnClick: true,
		Walkboxes:   scene.Floor,
		ExitLook:    script.ExitLook,
		ExitTake:    script.ExitTake,
		UseWithFail: script.UseWithFail,

		AchievementManager: achievement.AchievementManager,
		CharacterManager:   character.CharacterManager,
		DialogManager:      dialog.DialogManager,
		ItemManager:        item.ItemManager,
		SceneManager:       scene.SceneManager,
		ScriptManager:      script.ScriptManager,
		SpeakerManager:     speaker.SpeakerManager,
		StageManager:       stage.StageManager,
		ThemeManager:       theme.ThemeManager,
		WidgetManager:      widget.WidgetManager,
	})
}
