package scene

import ()

var PublicBBS = SceneManager.Register(&Scene{
	Name: "public_bbs",
	Title: Text{
		{Lang: "en", Text: "PUBLIC BBS TERMINAL"},
		{Lang: "hu", Text: "NYILVÁNOS BBS TERMINÁL"},
	},
})

var _ = PublicBBS.Leads(
	exitToSelector,
)
