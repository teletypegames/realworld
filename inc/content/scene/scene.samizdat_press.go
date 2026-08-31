package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name: inc.SceneSamizdatPress,
		Title: inkwell.Text{
			{Lang: "en", Text: "SAMIZDAT PRINTING HOUSE"},
			{Lang: "hu", Text: "SZAMIZDAT NYOMDA"},
		},
		Background: inc.BgSamizdatPress,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
