package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	SceneManager.Register(Scene{
		Name:       SceneHackerspace,
		Title:      "HACKERSPACE",
		Background: BgHackerspace,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
