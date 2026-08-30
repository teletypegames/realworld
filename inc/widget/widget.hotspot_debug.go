package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	Manager.Register(&inkwell.HotspotDebug{
		Name: "hotspot_debug",
	})
}
