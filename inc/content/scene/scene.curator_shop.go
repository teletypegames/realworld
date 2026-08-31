package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	SceneManager.Register(Scene{
		Name: inc.SceneCuratorShop,
		Title: inkwell.Text{
			{Lang: "en", Text: "CURATOR SERVICE — THE MASTER'S WORKSHOP"},
			{Lang: "hu", Text: "KURÁTORSZOLGÁLAT — A MESTER MŰHELYE"},
		},
		Background: inc.BackgroundCuratorShop,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
