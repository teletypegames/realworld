package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Background = inkwell.Asset

var BackgroundManager = NewManager(func(entity Background) string { return entity.Name })
