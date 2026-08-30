package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	WidgetManager.Register(gated("status", &inkwell.StatusLine{
		Name:        "status",
		Y:           statusY,
		Align:       inkwell.AlignLeft,
		ScreenWidth: DividerX,
	}))
}
