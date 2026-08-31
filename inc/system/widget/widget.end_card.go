package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	Manager.Register(&inkwell.EndCard{
		Name: "endcard",
	})
}
