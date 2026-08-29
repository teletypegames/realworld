package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	SceneManager.Register(Scene{
		Name:       SceneSamizdatPress,
		Title:      "SAMIZDAT PRINTING HOUSE",
		Background: BgSamizdatPress,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
