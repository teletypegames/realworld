package achievement

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type (
	Achievement = inkwell.Achievement
	Text        = inkwell.Text
)

var AchievementManager = inkwell.NewManager[*Achievement]()
