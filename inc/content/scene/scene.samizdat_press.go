package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

var SamizdatPress = SceneManager.Register(&Scene{
	Name: "samizdat_press",
	Title: inkwell.Text{
		{Lang: "en", Text: "SAMIZDAT PRINTING HOUSE"},
		{Lang: "hu", Text: "SZAMIZDAT NYOMDA"},
	},
})

var _ = SamizdatPress.Leads(
	exitToSelector,
)
