package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Item = inkwell.Item

var ItemManager = inkwell.NewManager[Item]()
