package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/widget"
)

var mockCombineOperator = Text{
	{Lang: "en", Text: "+"},
	{Lang: "hu", Text: "+"},
}

var MockCombine = mockChrome(StageManager.Register(&Stage{
	Name: "mock_combine",
	Title: Text{
		{Lang: "en", Text: "MOCK — COMBINING TWO ITEMS"},
		{Lang: "hu", Text: "MAKETT — KÉT TÁRGY ÖSSZERAKÁSA"},
	},
}))

var _ = widget.WidgetManager.Register(&widget.ArtMock{
	Name: "mock_combine_art",
	When: inkwell.OnStage(MockCombine),
	Dim:  true,
	Caption: Text{
		{Lang: "en", Text: "ALLEY — paused"},
		{Lang: "hu", Text: "SIKÁTOR — szünet"},
	},
})

var _ = widget.WidgetManager.Register(&widget.CombinePane{
	Name: "mock_combine_pane",
	When: inkwell.OnStage(MockCombine),
	Header: Text{
		{Lang: "en", Text: "TWO THINGS IN HAND"},
		{Lang: "hu", Text: "KÉT DOLOG A KÉZBEN"},
	},
	Hint: Text{
		{Lang: "en", Text: "chosen from the inventory, in order"},
		{Lang: "hu", Text: "a leltárból választva, sorban"},
	},
	Cards: []Text{
		{
			{Lang: "en", Text: "trinket"},
			{Lang: "hu", Text: "csecsebecse"},
		},
		{
			{Lang: "en", Text: "black Armilla"},
			{Lang: "hu", Text: "fekete Armilla"},
		},
	},
	Operator: mockCombineOperator,
	Caption: Text{
		{Lang: "en", Text: "WHAT YOU COULD DO"},
		{Lang: "hu", Text: "MIT TEHETNÉL"},
	},
	Rows: []widget.Choice{
		{
			Label: Text{
				{Lang: "en", Text: "Work it into the casing"},
				{Lang: "hu", Text: "Feszítsd bele a burkolatba"},
			},
		},
		{
			Label: Text{
				{Lang: "en", Text: "Scrape the contacts"},
				{Lang: "hu", Text: "Kapard le az érintkezőket"},
			},
		},
		{
			Label: Text{
				{Lang: "en", Text: "Never mind"},
				{Lang: "hu", Text: "Hagyd"},
			},
			Dim: true,
		},
	},
	Footer: Text{
		{Lang: "en", Text: "a wrong pairing just fails in character — trying costs nothing"},
		{Lang: "hu", Text: "a rossz párosítás csak jellemhűen elbukik — próbálni ingyen van"},
	},
})

var _ = widget.WidgetManager.Register(&widget.Console{
	Name:   "mock_combine_console",
	When:   inkwell.OnStage(MockCombine),
	Tokens: 3,
	Lines: []widget.ConsoleLine{
		{
			Body: Text{
				{Lang: "en", Text: "Trinket into the casing — that I'd try. Mind the contacts."},
				{Lang: "hu", Text: "A csecsebecsét a burkolatba — ezt megpróbálnám. Vigyázz az érintkezőkre."},
			},
			Cost: Text{
				{Lang: "en", Text: "— free: you proposed it"},
				{Lang: "hu", Text: "— ingyen: te vetetted fel"},
			},
		},
	},
})
