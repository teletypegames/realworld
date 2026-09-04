package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type (
	Scene      = inkwell.Scene
	Text       = inkwell.Text
	Action     = inkwell.Action
	Given      = inkwell.Given
	Hotspot    = inkwell.Hotspot
	Exit       = inkwell.Exit
	Trigger    = inkwell.Trigger
	SceneActor = inkwell.SceneActor
	Point      = inkwell.Point
	Polygon    = inkwell.Polygon
)

var SceneManager = inkwell.NewManager[*Scene]()

var Floor = []Polygon{
	inkwell.PolyRect(16, 196, 608, 62),
}
