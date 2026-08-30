package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.SceneHackerspace,
		Title:      "HACKERSPACE",
		Background: inc.BgHackerspace,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
