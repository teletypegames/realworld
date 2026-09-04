package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/background"
)

var _ = StageManager.Register(Stage{
	Name: inc.StageSplashStudio,
	Title: inkwell.Text{
		{Lang: "en", Text: "TELETYPE GAMES"},
		{Lang: "hu", Text: "TELETYPE GAMES"},
	},
	Background: inc.StageSplashStudio,
	Transition: splashTransition,
	Hold:       splashHold,
	Next:       inc.StageSplashTitle,
})

var _ = background.BackgroundManager.Register(background.Background{
	Name: inc.StageSplashStudio,
	Path: "assets/stage/splash_studio.png",
	Kind: inkwell.AssetImage,
})
