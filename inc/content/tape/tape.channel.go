package tape

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/speaker"
	"realworld/inc/system/widget"
)

var _ = widget.WidgetManager.Register(&inkwell.ChatLog{
	Name:       "channel",
	When:       widget.WhenPlaying,
	Bounds:     inkwell.Rect(widget.DividerX+2, widget.HUDTop+2, widget.ScreenWidth-widget.DividerX-4, widget.ScreenHeight-widget.HUDTop-4),
	LineHeight: widget.LineHeight,
	Padding:    widget.Padding,
	Speaker:    inkwell.SpeakerHeader,
	Deliveries: []inkwell.Delivery{inkwell.DeliveryLog},
	Resting:    speaker.Dex,
})
