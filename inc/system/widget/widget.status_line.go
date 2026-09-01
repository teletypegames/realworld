package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	WidgetManager.Register(&inkwell.StatusLine{
		Name:        "status",
		When:        WhenPlaying,
		Y:           statusY,
		Align:       inkwell.AlignLeft,
		ScreenWidth: DividerX,
	})
}
