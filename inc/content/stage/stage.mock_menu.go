package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/widget"
)

var MockMenu = mockChrome(StageManager.Register(&Stage{
	Name: "mock_menu",
	Title: Text{
		{Lang: "en", Text: "MOCK — MAIN MENU"},
		{Lang: "hu", Text: "MAKETT — FŐMENÜ"},
	},
}))

var _ = widget.WidgetManager.Register(&widget.ArtMock{
	Name: "mock_menu_art",
	When: inkwell.OnStage(MockMenu),
	Boxes: []widget.ArtBox{
		{
			Area: inkwell.Rect(170, 79, 300, 70),
			Label: Text{
				{Lang: "en", Text: "[ UNDERWATER SUN RISE ]"},
				{Lang: "hu", Text: "[ UNDERWATER SUN RISE ]"},
			},
		},
	},
	Subtitle: Text{
		{Lang: "en", Text: "attract-mode scene loops behind"},
		{Lang: "hu", Text: "a bemutató jelenet fut mögötte"},
	},
})

var _ = widget.WidgetManager.Register(&widget.ChoicePane{
	Name: "mock_menu_choices",
	When: inkwell.OnStage(MockMenu),
	Header: Text{
		{Lang: "en", Text: "SELECT"},
		{Lang: "hu", Text: "VÁLASSZ"},
	},
	Hint: Text{
		{Lang: "en", Text: "arrow keys or click — nothing here is typed"},
		{Lang: "hu", Text: "nyilak vagy kattintás — itt semmit sem kell begépelni"},
	},
	Rows: []widget.Choice{
		{
			Label: Text{
				{Lang: "en", Text: "Continue — the alley behind Noodle's house"},
				{Lang: "hu", Text: "Folytatás — a sikátor Noodle háza mögött"},
			},
		},
		{
			Label: Text{
				{Lang: "en", Text: "Restart the night"},
				{Lang: "hu", Text: "Kezdd újra az éjszakát"},
			},
		},
		{
			Label: Text{
				{Lang: "en", Text: "Load a saved night"},
				{Lang: "hu", Text: "Tölts be egy mentett éjszakát"},
			},
		},
		{
			Label: Text{
				{Lang: "en", Text: "Settings"},
				{Lang: "hu", Text: "Beállítások"},
			},
		},
		{
			Label: Text{
				{Lang: "en", Text: "Quit"},
				{Lang: "hu", Text: "Kilépés"},
			},
			Dim: true,
		},
	},
})

var _ = widget.WidgetManager.Register(&widget.Console{
	Name:   "mock_menu_console",
	When:   inkwell.OnStage(MockMenu),
	Tokens: 4,
	Lines: []widget.ConsoleLine{
		{
			Body: Text{
				{Lang: "en", Text: "Evening. You left off in an alley, holding an unlabelled tape."},
				{Lang: "hu", Text: "Este van. Egy sikátorban hagytad abba, címkézetlen kazettával."},
			},
		},
		{
			Body: Text{
				{Lang: "en", Text: "I can't open doors for you — only notice things."},
				{Lang: "hu", Text: "Nem tudok ajtót nyitni neked — csak észrevenni dolgokat."},
			},
			Cost: Text{
				{Lang: "en", Text: "— free"},
				{Lang: "hu", Text: "— ingyen"},
			},
		},
	},
})
