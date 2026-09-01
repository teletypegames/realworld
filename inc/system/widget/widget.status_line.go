package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

var _ = WidgetManager.Register(&inkwell.StatusLine{
	Name:        "status",
	When:        WhenPlaying,
	Y:           statusY,
	Align:       inkwell.AlignLeft,
	ScreenWidth: DividerX,
})
