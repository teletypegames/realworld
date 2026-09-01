package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

var _ = WidgetManager.Register(&inkwell.HotspotDebug{
	Name: "hotspot_debug",
})
