package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

var CuratorShop = SceneManager.Register(&Scene{
	Name: "curator_shop",
	Title: inkwell.Text{
		{Lang: "en", Text: "CURATOR SERVICE — THE MASTER'S WORKSHOP"},
		{Lang: "hu", Text: "KURÁTORSZOLGÁLAT — A MESTER MŰHELYE"},
	},
})

var _ = CuratorShop.Leads(
	exitToSelector,
)
