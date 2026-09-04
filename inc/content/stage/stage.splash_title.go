package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/widget"
)

var _ = StageManager.Register(Stage{
	Name: inc.StageSplashTitle,
	Title: inkwell.Text{
		{Lang: "en", Text: "REAL WORLD"},
		{Lang: "hu", Text: "REAL WORLD"},
	},
	Transition: splashTransition,
	Hold:       splashHold,
	Next:       inc.StageMainMenu,
})

var _ = widget.WidgetManager.Register(&inkwell.TextCard{
	Name:     "splash_title",
	When:     inkwell.OnStage(inc.StageSplashTitle),
	Align:    inkwell.AlignCenter,
	Backdrop: true,
	Lines: []inkwell.Text{
		{
			{Lang: "en", Text: "Real World"},
			{Lang: "hu", Text: "Real World"},
		},
	},
})
