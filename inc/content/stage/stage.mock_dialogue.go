package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/widget"
)

var mockDialogueBubble = Text{
	{Lang: "en", Text: "“Nobody's been down that back door since the raid. And nobody's going to, while I'm sitting here.”"},
	{Lang: "hu", Text: "„Senki sem ment be azon a hátsó ajtón a razzia óta. És amíg én itt ülök, nem is fog.”"},
}

var mockDialogueHeader = Text{
	{Lang: "en", Text: "YOU SAY"},
	{Lang: "hu", Text: "EZT MONDOD"},
}

var mockDialogueHint = Text{
	{Lang: "en", Text: "Paul speaks — Dex never answers for you"},
	{Lang: "hu", Text: "Paul beszél — Dex sosem válaszol helyetted"},
}

var mockDialogueRows = []widget.Choice{
	{
		Label: Text{
			{Lang: "en", Text: "“Then I'll wait until your shift ends.”"},
			{Lang: "hu", Text: "„Akkor megvárom, míg letelik a műszakod.”"},
		},
	},
	{
		Label: Text{
			{Lang: "en", Text: "“Who told the police to look here?”"},
			{Lang: "hu", Text: "„Ki mondta a rendőrségnek, hogy itt keressenek?”"},
		},
	},
	{
		Label: Text{
			{Lang: "en", Text: "[ show him the unmarked Personal Tape ]"},
			{Lang: "hu", Text: "[ mutasd meg neki a címkézetlen Személyes Kazettát ]"},
		},
	},
	{
		Label: Text{
			{Lang: "en", Text: "Say nothing and walk on"},
			{Lang: "hu", Text: "Ne szólj semmit, menj tovább"},
		},
		Dim: true,
	},
}

var mockDialogueLines = []widget.ConsoleLine{
	{
		Body: Text{
			{Lang: "en", Text: "His shift ends at four. A fact, not advice."},
			{Lang: "hu", Text: "A műszakja négykor ér véget. Tény, nem tanács."},
		},
		Cost: Text{
			{Lang: "en", Text: "— 1 token"},
			{Lang: "hu", Text: "— 1 zseton"},
		},
	},
	{
		Body: Text{
			{Lang: "en", Text: "His Armilla is the wrong series for this district."},
			{Lang: "hu", Text: "Az Armillája rossz sorozatú ehhez a kerülethez."},
		},
		Cost: Text{
			{Lang: "en", Text: "— free — I noticed it"},
			{Lang: "hu", Text: "— ingyen — én vettem észre"},
		},
	},
}

func mockDialogueArt(name string, showing inkwell.Condition) *widget.ArtMock {
	return &widget.ArtMock{
		Name: name,
		When: showing,
		Caption: Text{
			{Lang: "en", Text: "ALLEY — BEHIND NOODLE'S HOUSE"},
			{Lang: "hu", Text: "SIKÁTOR — NOODLE HÁZA MÖGÖTT"},
		},
		Boxes: []widget.ArtBox{
			{
				Area: inkwell.Rect(51, 98, 70, 108),
				Label: Text{
					{Lang: "en", Text: "PAUL"},
					{Lang: "hu", Text: "PAUL"},
				},
			},
			{
				Area: inkwell.Rect(291, 63, 76, 143),
				Label: Text{
					{Lang: "en", Text: "NOODLE"},
					{Lang: "hu", Text: "NOODLE"},
				},
			},
		},
		Bubble:   mockDialogueBubble,
		BubbleAt: inkwell.Rect(32, 26, 237, 0),
	}
}

var MockDialogue = mockChrome(StageManager.Register(&Stage{
	Name: "mock_dialogue",
	Title: Text{
		{Lang: "en", Text: "MOCK — NPC DIALOGUE"},
		{Lang: "hu", Text: "MAKETT — PÁRBESZÉD"},
	},
}))

var _ = widget.WidgetManager.Register(mockDialogueArt("mock_dialogue_art", inkwell.OnStage(MockDialogue)))

var _ = widget.WidgetManager.Register(&widget.ChoicePane{
	Name:   "mock_dialogue_choices",
	When:   inkwell.OnStage(MockDialogue),
	Header: mockDialogueHeader,
	Hint:   mockDialogueHint,
	Rows:   mockDialogueRows,
})

var _ = widget.WidgetManager.Register(&widget.Console{
	Name:   "mock_dialogue_console",
	When:   inkwell.OnStage(MockDialogue),
	Tokens: 3,
	Lines:  mockDialogueLines,
})
