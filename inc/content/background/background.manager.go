package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Background = inkwell.Asset

var BackgroundManager = inkwell.NewManager[Background]()
