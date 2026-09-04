package sound

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Sound = inkwell.Sound

var SoundManager = inkwell.NewManager[*Sound]()
