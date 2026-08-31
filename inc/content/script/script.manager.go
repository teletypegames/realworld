package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Script = inkwell.Script

var Manager = inkwell.NewManager[Script]()
