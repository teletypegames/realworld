package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var TapeInsert = ScriptManager.Register(&Script{
	Name: "tape_insert",
	Actions: inkwell.Seq(
		character.Paul.Say(Text{
			{Lang: "en", Text: "Right. Let's see who you are."},
			{Lang: "hu", Text: "Jól van. Lássuk, ki vagy."},
		}),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "Clicked in. Spinning. A stranger's tape in your own Armilla, Paul."},
			{Lang: "hu", Text: "Bekattant. Forog. Egy idegen kazettája a saját Armilládban, Paul."},
		}),
	),
})
