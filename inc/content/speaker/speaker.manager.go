package speaker

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Speaker = inkwell.Speaker

var SpeakerManager = inkwell.NewManager[*Speaker]()
