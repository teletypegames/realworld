package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	SceneManager.Register(Scene{
		Name: inc.SceneStreet,
		Title: inkwell.Text{
			{Lang: "en", Text: "STREET"},
			{Lang: "hu", Text: "UTCA"},
		},
		Background: inc.BackgroundStreet,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
