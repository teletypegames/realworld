package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	Manager.Register(&inkwell.StatusLine{
		Name:        "status",
		When:        whenPlaying,
		Y:           statusY,
		Align:       inkwell.AlignLeft,
		ScreenWidth: DividerX,
	})
}
