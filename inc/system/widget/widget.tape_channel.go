package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = WidgetManager.Register(&inkwell.ChatLog{
	Name:       "channel",
	When:       WhenPlaying,
	Bounds:     inkwell.Rect(DividerX+2, HUDTop+2, ScreenWidth-DividerX-4, ScreenHeight-HUDTop-4),
	LineHeight: LineHeight,
	Padding:    Padding,
	Speaker:    inkwell.SpeakerHeader,
	Voices:     []inkwell.Voice{inkwell.VoiceLog},
	Resting:    inc.Dex,
})
