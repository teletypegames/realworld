package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Character = inkwell.Character

var CharacterManager = inkwell.NewManager[*Character]()
