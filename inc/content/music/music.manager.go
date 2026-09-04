package music

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Music = inkwell.Music

var MusicManager = inkwell.NewManager[*Music]()
