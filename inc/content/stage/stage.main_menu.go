package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/widget"
)

var _ = StageManager.Register(Stage{
	Name: inc.StageMainMenu,
	Title: inkwell.Text{
		{Lang: "en", Text: "REAL WORLD"},
		{Lang: "hu", Text: "REAL WORLD"},
	},
})

var _ = widget.WidgetManager.Register(&inkwell.Menu{
	Name:   "main_menu",
	When:   inkwell.OnStage(inc.StageMainMenu),
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
					inkwell.GoToStage(inc.StageGame),
				),
			},
			{
				Label: inkwell.Text{
					{Lang: "en", Text: "Resume"},
					{Lang: "hu", Text: "Folytatás"},
				},
				Enabled: inkwell.Started(),
				Do:      inkwell.GoToStage(inc.StageGame),
			},
			{
				Label: inkwell.Text{
					{Lang: "en", Text: "Load"},
					{Lang: "hu", Text: "Betöltés"},
				},
				Do: inkwell.GoToStage(inc.StageLoad),
			},
			{
				Label: inkwell.Text{
					{Lang: "en", Text: "Save"},
					{Lang: "hu", Text: "Mentés"},
				},
				Enabled: inkwell.Started(),
				Do:      inkwell.GoToStage(inc.StageSave),
			},
			{
				Label: inkwell.Text{
					{Lang: "en", Text: "Settings"},
					{Lang: "hu", Text: "Beállítások"},
				},
				Do: inkwell.GoToStage(inc.StageSettings),
			},
			{
				Label: inkwell.Text{
					{Lang: "en", Text: "Credits"},
					{Lang: "hu", Text: "Készítők"},
				},
				Do: inkwell.GoToStage(inc.StageCredits),
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
