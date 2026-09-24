package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/widget"
)

var MockRail = mockChrome(StageManager.Register(&Stage{
	Name: "mock_rail",
	Title: Text{
		{Lang: "en", Text: "MOCK — CONSOLE COLLAPSED"},
		{Lang: "hu", Text: "MAKETT — BEHÚZOTT KONZOL"},
	},
	OnEnter: inkwell.SetFlag(inc.FlagConsoleCollapsed),
}))

var _ = widget.WidgetManager.Register(mockDialogueArt("mock_rail_art", inkwell.OnStage(MockRail)))

var _ = widget.WidgetManager.Register(&widget.ChoicePane{
	Name:   "mock_rail_choices",
	When:   inkwell.OnStage(MockRail),
	Header: mockDialogueHeader,
	Hint:   mockDialogueHint,
	Rows:   mockDialogueRows,
})

var _ = widget.WidgetManager.Register(&widget.Console{
	Name:   "mock_rail_console",
	When:   inkwell.OnStage(MockRail),
	Tokens: 3,
	Lines:  mockDialogueLines,
})
