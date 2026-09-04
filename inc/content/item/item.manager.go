package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type (
	Item   = inkwell.Item
	Text   = inkwell.Text
	Action = inkwell.Action
)

var ItemManager = inkwell.NewManager[*Item]()
