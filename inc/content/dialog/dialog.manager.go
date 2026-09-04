package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Dialog = inkwell.Dialog

var DialogManager = inkwell.NewManager[*Dialog]()
