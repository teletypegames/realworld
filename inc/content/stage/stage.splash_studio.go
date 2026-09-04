package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/background"
)

var SplashStudio = StageManager.Register(&Stage{
	Name: "splash_studio",
	Title: inkwell.Text{
		{Lang: "en", Text: "TELETYPE GAMES"},
		{Lang: "hu", Text: "TELETYPE GAMES"},
	},
	Background: background.BackgroundManager.Register(&background.Background{
		Name: "splash_studio",
		Path: "assets/stage/splash_studio.png",
		Kind: inkwell.AssetImage,
	}),
	Transition: splashTransition,
	Hold:       splashHold,
	Next:       SplashTitle,
})
