package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/widget"
)

var MockCutscene = mockChrome(StageManager.Register(&Stage{
	Name: "mock_cutscene",
	Title: Text{
		{Lang: "en", Text: "MOCK — CUTSCENE"},
		{Lang: "hu", Text: "MAKETT — ÁTVEZETŐ"},
	},
	OnEnter: inkwell.SetFlag(inc.FlagUICollapsed),
}))

var _ = widget.WidgetManager.Register(&widget.ArtMock{
	Name: "mock_cutscene_art",
	When: inkwell.OnStage(MockCutscene),
	Boxes: []widget.ArtBox{
		{
			Area: inkwell.Rect(211, 45, 218, 138),
			Label: Text{
				{Lang: "en", Text: "[ CLOSE-UP — THE UNMARKED TAPE ]"},
				{Lang: "hu", Text: "[ KÖZELI — A CÍMKÉZETLEN KAZETTA ]"},
			},
		},
	},
})

var _ = widget.WidgetManager.Register(&widget.ChoicePane{
	Name: "mock_cutscene_choices",
	When: inkwell.OnStage(MockCutscene),
	Header: Text{
		{Lang: "en", Text: "CUTSCENE"},
		{Lang: "hu", Text: "ÁTVEZETŐ"},
	},
	Hint: Text{
		{Lang: "en", Text: "no choices here — any key to move on"},
		{Lang: "hu", Text: "itt nincs választás — bármelyik gomb továbblép"},
	},
	Speaker: Text{
		{Lang: "en", Text: "PAUL"},
		{Lang: "hu", Text: "PAUL"},
	},
	Body: Text{
		{Lang: "en", Text: "“No label, no seal. Somebody wound this one by hand.”"},
		{Lang: "hu", Text: "„Se címke, se pecsét. Ezt valaki kézzel tekerte fel.”"},
	},
	Footer: Text{
		{Lang: "en", Text: "SKIP →"},
		{Lang: "hu", Text: "ÁTUGRÁS →"},
	},
})

var _ = widget.WidgetManager.Register(&widget.Console{
	Name:   "mock_cutscene_console",
	When:   inkwell.OnStage(MockCutscene),
	Tokens: 3,
	Lines: []widget.ConsoleLine{
		{
			Body: Text{
				{Lang: "en", Text: "I've heard a hand-wound tape before."},
				{Lang: "hu", Text: "Hallottam már kézzel tekert kazettát."},
			},
		},
		{
			Body: Text{
				{Lang: "en", Text: "Hold TAB to hear where."},
				{Lang: "hu", Text: "Tartsd nyomva a TAB-ot, és elmondom, hol."},
			},
			Cost: Text{
				{Lang: "en", Text: "— 1 token"},
				{Lang: "hu", Text: "— 1 zseton"},
			},
		},
	},
})
