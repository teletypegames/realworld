package stage

var SplashStudio = StageManager.Register(&Stage{
	Name: "splash_studio",
	Title: Text{
		{Lang: "en", Text: "TELETYPE GAMES"},
		{Lang: "hu", Text: "TELETYPE GAMES"},
	},
	Background: &Asset{
		Name: "splash_studio",
		Path: "assets/stage/splash_studio.png",
	},
	Transition: splashTransition,
	Hold:       splashHold,
	Skippable:  true,
	Next:       SplashTitle,
})
