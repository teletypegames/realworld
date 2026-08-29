package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Dialog = inkwell.Dialogue

var DialogManager = NewManager(func(entity Dialog) string { return entity.Name })
