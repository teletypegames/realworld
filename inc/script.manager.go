package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Script = inkwell.Script

var ScriptManager = inkwell.NewManager[Script]()
