package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	WidgetManager.Register(&inkwell.SceneNav{
		Name: "scene_nav",
		When: whenUnpaused,
	})
}
