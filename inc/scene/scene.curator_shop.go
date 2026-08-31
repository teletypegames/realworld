package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name: inc.SceneCuratorShop,
		Title: inkwell.Text{
			{Lang: "en", Text: "CURATOR SERVICE — THE MASTER'S WORKSHOP"},
			{Lang: "hu", Text: "KURÁTORSZOLGÁLAT — A MESTER MŰHELYE"},
		},
		Background: inc.BgCuratorShop,
		Exits: []inkwell.Exit{
			exitToSelector,
		},
	})
}
