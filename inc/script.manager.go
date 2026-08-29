package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Script = inkwell.Script

var ScriptManager = NewManager(func(entity Script) string { return entity.Name })
