package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	Manager.Register(Script{
		Name: inc.ScriptOpening,
		Actions: inkwell.Seq(
			inkwell.SetFlag(inc.FlagPoliceTip),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "Back alley. Just like the clerk said."},
				{Lang: "hu", Text: "Hátsó sikátor. Pont, ahogy az ügyintéző mondta."},
			}),
			inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "A government office tipped you off to remove something from a scene. That doesn't strike you as odd?"},
				{Lang: "hu", Text: "Egy állami hivatal súgta meg, hogy vigyél el valamit egy helyszínről. Ez neked nem furcsa?"},
			}),
		),
	})
}
