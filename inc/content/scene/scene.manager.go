package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Scene = inkwell.Scene

var SceneManager = inkwell.NewManager[Scene]()

var Floor = []inkwell.Polygon{
	inkwell.Poly(
		inkwell.Point{
			X: 16,
			Y: 196,
		}, inkwell.Point{
			X: 624,
			Y: 196,
		},
		inkwell.Point{
			X: 624,
			Y: 258,
		}, inkwell.Point{
			X: 16,
			Y: 258,
		},
	),
}
