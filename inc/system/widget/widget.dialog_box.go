package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

var _ = WidgetManager.Register(&inkwell.DialogBox{
	Name:       "dialog",
	Bounds:     inkwell.Rect(0, ScreenHeight-LineHeight*9, DividerX, LineHeight*9),
	LineHeight: LineHeight,
	Padding:    Padding,
})
