package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type (
	Character = inkwell.Character
	Text      = inkwell.Text
	Point     = inkwell.Point
)

var CharacterManager = inkwell.NewManager[*Character]()
