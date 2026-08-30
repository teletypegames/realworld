package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	Manager.Register(&inkwell.SpeechBubble{
		Name:      "speech",
		MaxWidth:  360,
		Padding:   Pad,
		OffsetY:   8,
		FallbackY: TopBarH + Pad,
	})
}
