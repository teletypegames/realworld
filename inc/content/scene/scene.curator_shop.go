package scene

import ()

var CuratorShop = SceneManager.Register(&Scene{
	Name: "curator_shop",
	Title: Text{
		{Lang: "en", Text: "CURATOR SERVICE — THE MASTER'S WORKSHOP"},
		{Lang: "hu", Text: "KURÁTORSZOLGÁLAT — A MESTER MŰHELYE"},
	},
})

var _ = CuratorShop.Leads(
	exitToSelector,
)
