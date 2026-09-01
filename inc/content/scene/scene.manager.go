package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Scene = inkwell.Scene

var SceneManager = inkwell.NewManager[Scene]()

var Floor = []inkwell.Polygon{
	inkwell.PolyRect(16, 196, 608, 62),
}
