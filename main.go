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
	"realworld/inc/content/stage"
	"realworld/inc/content/tape"
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
	startScene string
	finale     bool
	language   string
	debug      bool
}

func parseOptions() options {
	startScene := flag.String("scene", "", "start straight in this scene, skipping the main menu")
	finale := flag.Bool("finale", false, "start with the finale, to try the Nokia-punk theme switch")
	language := flag.String("lang", "en", "language the game speaks (en, hu)")
	debug := flag.Bool("debug", false, "open the debug menu with d (a debug build has it anyway)")
	flag.Parse()

	return options{
		startScene: *startScene,
		finale:     *finale,
		language:   *language,
		debug:      *debug,
	}
}

// startStage is the main menu, unless the run named a scene or asked for the
// finale — both of those are somebody wanting to look at the game, not at its
// front door.
func (chosen options) startStage() string {
	if chosen.startScene != "" || chosen.finale {
		return inc.StageGame
	}
	return inc.StageMainMenu
}

func (chosen options) scene() string {
	if chosen.startScene == "" {
		return inc.ScenePaulShop
	}
	return chosen.startScene
}

func (chosen options) finaleScript() string {
	if !chosen.finale {
		return ""
	}
	return inc.ScriptFinale
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
		Width:       widget.ScreenWidth,
		Height:      widget.ScreenHeight,
		WindowScale: 2,
		MaxLogLines: 64,
		SceneRect:   inkwell.Rect(0, widget.TopBarHeight, widget.ScreenWidth, widget.HUDTop-widget.TopBarHeight),

		StartStage:   chosen.startStage(),
		StartScene:   chosen.scene(),
		StartScript:  inc.ScriptOpening,
		FinaleScript: chosen.finaleScript(),
		StartTheme:   inc.RealWorld,
		Flags:        chosen.flags(),

		Player:      inc.Paul,
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
		SpeakerManager:   tape.Speakers,
		StageManager:     stage.StageManager,
		ThemeManager:     theme.ThemeManager,
		WidgetManager:    widget.WidgetManager,
	})
}
