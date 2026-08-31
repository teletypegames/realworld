package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system"
)

func init() {
	ScriptManager.Register(Script{
		Name: inc.ScriptFinale,
		Actions: inkwell.Seq(
			system.System.SetMode(system.ModeCutscene),
			inkwell.Wait(0.6),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "I'm putting it in. That's all this is."},
				{Lang: "hu", Text: "Beteszem. Ennyi az egész."},
			}),
			inkwell.Wait(0.4),
			system.System.UseTheme(inc.NokiaPunk),
			inkwell.Say("norman", inkwell.Text{
				{Lang: "en", Text: "No — this isn't what you were supposed to do!"},
				{Lang: "hu", Text: "Ne — nem ezt kellett volna tenned!"},
			}),
			inkwell.Wait(0.6),
			inkwell.Say(inc.TapeMystery, inkwell.Text{
				{Lang: "en", Text: "Thank you, Paul. You did exactly what I asked, the whole way through."},
				{Lang: "hu", Text: "Köszönöm, Paul. Pontosan azt tetted, amit kértem, végig."},
			}),
			inkwell.Wait(1.0),
			inkwell.ShowEnd(inkwell.Text{
				{Lang: "en", Text: "REAL WORLD — end of game two"},
				{Lang: "hu", Text: "REAL WORLD — a második játék vége"},
			}),
		),
	})
}
