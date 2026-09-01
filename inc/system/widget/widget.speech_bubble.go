package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

var _ = WidgetManager.Register(&inkwell.SpeechBubble{
	Name:      "speech",
	MaxWidth:  360,
	Padding:   Padding,
	OffsetY:   8,
	FallbackY: TopBarHeight + Padding,
})
