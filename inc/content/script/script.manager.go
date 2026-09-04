package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type (
	Script = inkwell.Script
	Text   = inkwell.Text
	Action = inkwell.Action
	Ladder = inkwell.Ladder
)

var ScriptManager = inkwell.NewManager[*Script]()
