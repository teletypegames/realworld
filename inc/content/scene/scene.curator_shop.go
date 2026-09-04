package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type curatorShop struct{}

var _ place = curatorShop{}

var CuratorShop = SceneManager.Register(&Scene{
	Name: "curator_shop",
	Title: inkwell.Text{
		{Lang: "en", Text: "CURATOR SERVICE — THE MASTER'S WORKSHOP"},
		{Lang: "hu", Text: "KURÁTORSZOLGÁLAT — A MESTER MŰHELYE"},
	},
	Exits: []inkwell.Exit{
		exitToSelector,
	},
})

func (curatorShop) Scene() *inkwell.Scene { return CuratorShop }
