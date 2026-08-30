package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.SceneSamizdatPress,
		Title:      "SAMIZDAT PRINTING HOUSE",
		Background: inc.BgSamizdatPress,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
