package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/widget"
)

var Credits = StageManager.Register(&Stage{
	Name: "credits",
	Title: inkwell.Text{
		{Lang: "en", Text: "CREDITS"},
		{Lang: "hu", Text: "KÉSZÍTŐK"},
	},
})

var _ = widget.WidgetManager.Register(&inkwell.TextCard{
	Name:     "credits",
	When:     inkwell.OnStage(Credits),
	Align:    inkwell.AlignCenter,
	Backdrop: true,
	Lines: []inkwell.Text{
		{
			{Lang: "en", Text: "REAL WORLD"},
			{Lang: "hu", Text: "REAL WORLD"},
		},
		{},
		{
			{Lang: "en", Text: "Game two of the Norman Arc"},
			{Lang: "hu", Text: "A Norman-ív második játéka"},
		},
		{},
		{
			{Lang: "en", Text: "Teletype Games"},
			{Lang: "hu", Text: "Teletype Games"},
		},
		{
			{Lang: "en", Text: "Built on the inkwell engine"},
			{Lang: "hu", Text: "Az inkwell motorra épült"},
		},
		{},
		{},
		{
			{Lang: "en", Text: "ESC — back"},
			{Lang: "hu", Text: "ESC — vissza"},
		},
	},
})

var _ = widget.WidgetManager.Register(escapeToMenu("credits_escape", Credits))
