package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Item = inkwell.Item

var ItemManager = NewManager(func(entity Item) string { return entity.Name })
