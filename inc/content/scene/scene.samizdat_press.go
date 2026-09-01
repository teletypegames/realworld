package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	SceneManager.Register(Scene{
		Name: inc.SceneSamizdatPress,
		Title: inkwell.Text{
			{Lang: "en", Text: "SAMIZDAT PRINTING HOUSE"},
			{Lang: "hu", Text: "SZAMIZDAT NYOMDA"},
		},
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
