package theme

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Theme = inkwell.Theme

var ThemeManager = inkwell.NewManager[*Theme]()
