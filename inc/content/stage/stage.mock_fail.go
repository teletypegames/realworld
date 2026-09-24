package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/widget"
)

var MockFail = mockChrome(StageManager.Register(&Stage{
	Name: "mock_fail",
	Title: Text{
		{Lang: "en", Text: "MOCK — A COMBINE THAT FAILS"},
		{Lang: "hu", Text: "MAKETT — SIKERTELEN ÖSSZERAKÁS"},
	},
}))

var _ = widget.WidgetManager.Register(&widget.ArtMock{
	Name: "mock_fail_art",
	When: inkwell.OnStage(MockFail),
	Dim:  true,
	Caption: Text{
		{Lang: "en", Text: "ALLEY — paused"},
		{Lang: "hu", Text: "SIKÁTOR — szünet"},
	},
	Subtitle: Text{
		{Lang: "en", Text: "“…no. That's not going to read.”"},
		{Lang: "hu", Text: "„…nem. Ezt nem fogja beolvasni.”"},
	},
})

var _ = widget.WidgetManager.Register(&widget.CombinePane{
	Name: "mock_fail_pane",
	When: inkwell.OnStage(MockFail),
	Header: Text{
		{Lang: "en", Text: "IT DIDN'T TAKE"},
		{Lang: "hu", Text: "NEM JÖTT ÖSSZE"},
	},
	Hint: Text{
		{Lang: "en", Text: "both things go back in your coat, unharmed"},
		{Lang: "hu", Text: "mindkettő sértetlenül visszakerül a kabátodba"},
	},
	Cards: []Text{
		{
			{Lang: "en", Text: "unmarked tape"},
			{Lang: "hu", Text: "címkézetlen kazetta"},
		},
		{
			{Lang: "en", Text: "Noodle's letter"},
			{Lang: "hu", Text: "Noodle levele"},
		},
	},
	Operator: mockCombineOperator,
	Dimmed:   true,
	Rows: []widget.Choice{
		{
			Label: Text{
				{Lang: "en", Text: "Wind it round the spool"},
				{Lang: "hu", Text: "Tekerd rá az orsóra"},
			},
			Struck: true,
		},
		{
			Label: Text{
				{Lang: "en", Text: "Read the letter again"},
				{Lang: "hu", Text: "Olvasd el újra a levelet"},
			},
		},
	},
	TriedTop: Text{
		{Lang: "en", Text: "TRIED TONIGHT — 3"},
		{Lang: "hu", Text: "MA ESTE MEGPRÓBÁLVA — 3"},
	},
	Tried: []widget.TriedRow{
		{
			Pair: Text{
				{Lang: "en", Text: "trinket + letter"},
				{Lang: "hu", Text: "csecsebecse + levél"},
			},
			Quip: Text{
				{Lang: "en", Text: "“Now the letter has a hole in it.”"},
				{Lang: "hu", Text: "„Most már lyukas a levél.”"},
			},
		},
		{
			Pair: Text{
				{Lang: "en", Text: "tape + black Armilla"},
				{Lang: "hu", Text: "kazetta + fekete Armilla"},
			},
			Quip: Text{
				{Lang: "en", Text: "“It asked for a password. In an advert.”"},
				{Lang: "hu", Text: "„Jelszót kért. Egy reklámban.”"},
			},
		},
	},
})

var _ = widget.WidgetManager.Register(&widget.Console{
	Name:   "mock_fail_console",
	When:   inkwell.OnStage(MockFail),
	Tokens: 3,
	Lines: []widget.ConsoleLine{
		{
			Body: Text{
				{Lang: "en", Text: "You wound a letter onto a locked spool. Twice."},
				{Lang: "hu", Text: "Egy levelet tekertél egy zárt orsóra. Kétszer."},
			},
		},
		{
			Body: Text{
				{Lang: "en", Text: "The spool remains locked. So, largely, does the alley."},
				{Lang: "hu", Text: "Az orsó továbbra is zárva. Ahogy nagyrészt a sikátor is."},
			},
			Cost: Text{
				{Lang: "en", Text: "0 tokens — schadenfreude is free"},
				{Lang: "hu", Text: "0 zseton — a káröröm ingyen van"},
			},
		},
	},
})
