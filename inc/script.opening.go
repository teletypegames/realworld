package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	ScriptManager.Register(Script{
		Name: ScriptOpening,
		Actions: inkwell.Seq(
			inkwell.SetFlag(FlagPoliceTip),
			inkwell.Say(Paul, "Back alley. Just like the clerk said."),
			inkwell.Say(Dex, "A government office tipped you off to remove something from a scene. That doesn't strike you as odd?"),
		),
	})
}
