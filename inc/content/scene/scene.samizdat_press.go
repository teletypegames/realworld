package scene

import ()

var SamizdatPress = SceneManager.Register(&Scene{
	Name: "samizdat_press",
	Title: Text{
		{Lang: "en", Text: "SAMIZDAT PRINTING HOUSE"},
		{Lang: "hu", Text: "SZAMIZDAT NYOMDA"},
	},
})

var _ = SamizdatPress.Leads(
	exitToSelector,
)
