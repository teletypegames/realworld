package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/widget"
)

var MainMenu = StageManager.Register(&Stage{
	Name: "main_menu",
	Title: inkwell.Text{
		{Lang: "en", Text: "REAL WORLD"},
		{Lang: "hu", Text: "REAL WORLD"},
	},
})

var _ = widget.WidgetManager.Register(&inkwell.Menu{
	Name:   "main_menu",
	When:   inkwell.OnStage(MainMenu),
	Sticky: true,
	Bounds: window(240, 200),
	Root: inkwell.MenuItem{
		Label: inkwell.Text{
			{Lang: "en", Text: "REAL WORLD"},
			{Lang: "hu", Text: "REAL WORLD"},
		},
		Items: []inkwell.MenuItem{
			{
				Label: inkwell.Text{
					{Lang: "en", Text: "New Game"},
					{Lang: "hu", Text: "Új játék"},
				},
				Do: inkwell.Seq(
					inkwell.Restart(),
					inkwell.GoToStage(Game),
				),
			},
			{
				Label: inkwell.Text{
					{Lang: "en", Text: "Resume"},
					{Lang: "hu", Text: "Folytatás"},
				},
				Enabled: inkwell.Started(),
				Do:      inkwell.GoToStage(Game),
			},
			{
				Label: inkwell.Text{
					{Lang: "en", Text: "Load"},
					{Lang: "hu", Text: "Betöltés"},
				},
				Do: inkwell.GoToStage(Load),
			},
			{
				Label: inkwell.Text{
					{Lang: "en", Text: "Save"},
					{Lang: "hu", Text: "Mentés"},
				},
				Enabled: inkwell.Started(),
				Do:      inkwell.GoToStage(Save),
			},
			{
				Label: inkwell.Text{
					{Lang: "en", Text: "Settings"},
					{Lang: "hu", Text: "Beállítások"},
				},
				Do: inkwell.GoToStage(Settings),
			},
			{
				Label: inkwell.Text{
					{Lang: "en", Text: "Credits"},
					{Lang: "hu", Text: "Készítők"},
				},
				Do: inkwell.GoToStage(Credits),
			},
			{
				Label: inkwell.Text{
					{Lang: "en", Text: "Exit"},
					{Lang: "hu", Text: "Kilépés"},
				},
				Do: inkwell.Quit(),
			},
		},
	},
})
