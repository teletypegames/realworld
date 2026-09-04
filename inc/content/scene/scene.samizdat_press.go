package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type samizdatPress struct{}

var SamizdatPress = SceneManager.Register(&Scene{
	Name: "samizdat_press",
	Title: inkwell.Text{
		{Lang: "en", Text: "SAMIZDAT PRINTING HOUSE"},
		{Lang: "hu", Text: "SZAMIZDAT NYOMDA"},
	},
	Exits: []inkwell.Exit{
		exitToSelector,
	},
})

func (samizdatPress) Scene() *inkwell.Scene { return SamizdatPress }
