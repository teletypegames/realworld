package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Dialog = inkwell.Dialogue

var Manager = inkwell.NewManager[Dialog]()
