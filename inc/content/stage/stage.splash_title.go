package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/widget"
)

var SplashTitle = StageManager.Register(&Stage{
	Name: "splash_title",
	Title: Text{
		{Lang: "en", Text: "REAL WORLD"},
		{Lang: "hu", Text: "REAL WORLD"},
	},
	Transition: splashTransition,
	Hold:       splashHold,
	Skippable:  true,
	Next:       MainMenu,
})

var _ = widget.WidgetManager.Register(&inkwell.TextCard{
	Name:     "splash_title",
	When:     inkwell.OnStage(SplashTitle),
	Align:    inkwell.AlignCenter,
	Backdrop: true,
	Lines: []Text{
		{
			{Lang: "en", Text: "Underwater Sun Rising"},
			{Lang: "hu", Text: "Underwater Sun Rising"},
		},
	},
})
