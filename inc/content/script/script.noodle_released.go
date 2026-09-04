package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var NoodleReleased = ScriptManager.Register(&Script{
	Name: "noodle_released",
	Actions: inkwell.Seq(
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "Something on the police band, Paul. They let the antiques dealer go. No charge, no explanation."},
			{Lang: "hu", Text: "Valami a rendőrségi sávon, Paul. Elengedték a régiségkereskedőt. Nincs vád, nincs magyarázat."},
		}),
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "Noodle's out. He'll be at his office, then. That door's open now."},
			{Lang: "hu", Text: "Noodle kint van. Akkor az irodájában lesz. Az az ajtó most nyitva."},
		}),
		inkwell.SetFlag(inc.FlagNoodleReleased),
	),
})
