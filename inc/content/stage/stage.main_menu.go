package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/widget"
)

var MainMenu = StageManager.Register(&Stage{
	Name: "main_menu",
	Title: Text{
		{Lang: "en", Text: "REAL WORLD"},
		{Lang: "hu", Text: "REAL WORLD"},
	},
})

var _ = widget.WidgetManager.Register(&inkwell.Menu{
	Name:   "main_menu",
	When:   inkwell.OnStage(MainMenu),
	Sticky: true,
	Bounds: window(240, 200),
	Root: MenuItem{
		Label: Text{
			{Lang: "en", Text: "REAL WORLD"},
			{Lang: "hu", Text: "REAL WORLD"},
		},
		Items: []MenuItem{
			{
				Label: Text{
					{Lang: "en", Text: "New Game"},
					{Lang: "hu", Text: "Új játék"},
				},
				Do: inkwell.Seq(
					inkwell.Restart(),
					inkwell.GoToStage(Game),
				),
			},
			{
				Label: Text{
					{Lang: "en", Text: "Resume"},
					{Lang: "hu", Text: "Folytatás"},
				},
				Enabled: inkwell.Started(),
				Do:      inkwell.GoToStage(Game),
			},
			{
				Label: Text{
					{Lang: "en", Text: "Load"},
					{Lang: "hu", Text: "Betöltés"},
				},
				Do: inkwell.GoToStage(Load),
			},
			{
				Label: Text{
					{Lang: "en", Text: "Save"},
					{Lang: "hu", Text: "Mentés"},
				},
				Enabled: inkwell.Started(),
				Do:      inkwell.GoToStage(Save),
			},
			{
				Label: Text{
					{Lang: "en", Text: "Settings"},
					{Lang: "hu", Text: "Beállítások"},
				},
				Do: inkwell.GoToStage(Settings),
			},
			{
				Label: Text{
					{Lang: "en", Text: "Credits"},
					{Lang: "hu", Text: "Készítők"},
				},
				Do: inkwell.GoToStage(Credits),
			},
			{
				Label: Text{
					{Lang: "en", Text: "Exit"},
					{Lang: "hu", Text: "Kilépés"},
				},
				Do: inkwell.Quit(),
			},
		},
	},
})
