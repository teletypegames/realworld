package speaker

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type (
	Speaker = inkwell.Speaker
	Text    = inkwell.Text
)

var SpeakerManager = inkwell.NewManager[*Speaker]()
