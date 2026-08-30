package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Item = inkwell.Item

var Manager = inkwell.NewManager[Item]()
