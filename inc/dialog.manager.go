package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Dialog = inkwell.Dialogue

var DialogManager = inkwell.NewManager[Dialog]()
