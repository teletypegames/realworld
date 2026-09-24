package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/widget"
)

var MockUse = mockChrome(StageManager.Register(&Stage{
	Name: "mock_use",
	Title: Text{
		{Lang: "en", Text: "MOCK — USING AN ITEM"},
		{Lang: "hu", Text: "MAKETT — TÁRGY HASZNÁLATA"},
	},
}))

var _ = widget.WidgetManager.Register(&widget.ArtMock{
	Name: "mock_use_art",
	When: inkwell.OnStage(MockUse),
	Caption: Text{
		{Lang: "en", Text: "ALLEY — BEHIND NOODLE'S HOUSE"},
		{Lang: "hu", Text: "SIKÁTOR — NOODLE HÁZA MÖGÖTT"},
	},
	Boxes: []widget.ArtBox{
		{
			Area: inkwell.Rect(41, 81, 67, 125),
			Label: Text{
				{Lang: "en", Text: "PAUL"},
				{Lang: "hu", Text: "PAUL"},
			},
		},
		{
			Area:  inkwell.Rect(285, 57, 156, 109),
			Solid: true,
			Label: Text{
				{Lang: "en", Text: "gap at the foot of the wall"},
				{Lang: "hu", Text: "rés a fal tövében"},
			},
		},
	},
	Tag: Text{
		{Lang: "en", Text: "holding: trinket"},
		{Lang: "hu", Text: "kézben: csecsebecse"},
	},
	TagAt: inkwell.Point{
		X: 419,
		Y: 143,
	},
})

var _ = widget.WidgetManager.Register(&widget.ChoicePane{
	Name: "mock_use_choices",
	When: inkwell.OnStage(MockUse),
	Header: Text{
		{Lang: "en", Text: "WITH THE GAP"},
		{Lang: "hu", Text: "A RÉSSEL"},
	},
	Hint: Text{
		{Lang: "en", Text: "Paul does it — Dex cannot"},
		{Lang: "hu", Text: "Paul csinálja — Dex nem tudja"},
	},
	Rows: []widget.Choice{
		{
			Label: Text{
				{Lang: "en", Text: "Prise the loose brick out with the trinket"},
				{Lang: "hu", Text: "Feszítsd ki a laza téglát a csecsebecsével"},
			},
		},
		{
			Label: Text{
				{Lang: "en", Text: "Look into the gap"},
				{Lang: "hu", Text: "Nézz bele a résbe"},
			},
		},
		{
			Label: Text{
				{Lang: "en", Text: "Put the trinket away"},
				{Lang: "hu", Text: "Tedd el a csecsebecsét"},
			},
		},
		{
			Label: Text{
				{Lang: "en", Text: "Step back"},
				{Lang: "hu", Text: "Lépj hátrébb"},
			},
			Dim: true,
		},
	},
})

var _ = widget.WidgetManager.Register(&widget.Console{
	Name:   "mock_use_console",
	When:   inkwell.OnStage(MockUse),
	Tokens: 3,
	Lines: []widget.ConsoleLine{
		{
			Body: Text{
				{Lang: "en", Text: "The trinket is softer than the mortar. You have a wire somewhere."},
				{Lang: "hu", Text: "A csecsebecse puhább a habarcsnál. Van nálad valahol egy drót."},
			},
			Cost: Text{
				{Lang: "en", Text: "— 1 token"},
				{Lang: "hu", Text: "— 1 zseton"},
			},
		},
		{
			Body: Text{
				{Lang: "en", Text: "I'd not lever it yet. But it's your arm."},
				{Lang: "hu", Text: "Én még nem feszíteném. De a te karod."},
			},
			Cost: Text{
				{Lang: "en", Text: "— free"},
				{Lang: "hu", Text: "— ingyen"},
			},
		},
	},
})
