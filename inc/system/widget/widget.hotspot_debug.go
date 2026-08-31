package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	WidgetManager.Register(&inkwell.HotspotDebug{
		Name: "hotspot_debug",
	})
}
