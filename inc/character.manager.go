package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Character = inkwell.Character

var CharacterManager = NewManager(func(entity Character) string { return entity.Name })
