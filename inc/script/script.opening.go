package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Script{
		Name: inc.ScriptOpening,
		Actions: inkwell.Seq(
			inkwell.SetFlag(inc.FlagPoliceTip),
			inkwell.Say(inc.Paul, "Back alley. Just like the clerk said."),
			inkwell.Say(inc.Dex, "A government office tipped you off to remove something from a scene. That doesn't strike you as odd?"),
		),
	})
}
