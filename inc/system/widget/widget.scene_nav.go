package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	Manager.Register(&inkwell.SceneNav{
		Name: "scene_nav",
		When: whenUnpaused,
	})
}
