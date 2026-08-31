package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	WidgetManager.Register(&inkwell.DialogBox{
		Name:       "dialog",
		Bounds:     inkwell.Rect(0, ScreenH-LineH*9, DividerX, LineH*9),
		LineHeight: LineH,
		Padding:    Pad,
	})
}
