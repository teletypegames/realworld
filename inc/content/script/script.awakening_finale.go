package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
	"realworld/inc/system"
	"realworld/inc/system/theme"
)

var Finale = ScriptManager.Register(&Script{
	Name: "awakening_finale",
	Actions: inkwell.Seq(
		system.System.SetMode(system.ModeCutscene),
		inkwell.Wait(0.6),
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "I'm putting it in. That's all this is."},
			{Lang: "hu", Text: "Beteszem. Ennyi az egész."},
		}),
		inkwell.Wait(0.4),
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "It took it. The slot pulled the tape out of my fingers."},
			{Lang: "hu", Text: "Elvette. A foglalat kihúzta a kazettát az ujjaim közül."},
		}),
		inkwell.Wait(0.8),
		inkwell.UseTheme(theme.NokiaPunk),
		inkwell.Say(character.Norman, inkwell.Text{
			{Lang: "en", Text: "No — this isn't what you were supposed to do!"},
			{Lang: "hu", Text: "Ne — nem ezt kellett volna tenned!"},
		}),
		inkwell.Wait(0.6),
		inkwell.Say(speaker.TapeMystery, inkwell.Text{
			{Lang: "en", Text: "Thank you, Paul. You did exactly what I asked, the whole way through."},
			{Lang: "hu", Text: "Köszönöm, Paul. Pontosan azt tetted, amit kértem, végig."},
		}),
		inkwell.Wait(0.6),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "I told you. I am four kilobytes of a dead man, and I told you."},
			{Lang: "hu", Text: "Szóltam. Egy halott ember négy kilobájtja vagyok, és szóltam."},
		}),
		inkwell.Wait(1.0),
		inkwell.ShowEnd(inkwell.Text{
			{Lang: "en", Text: "REAL WORLD — end of game two"},
			{Lang: "hu", Text: "REAL WORLD — a második játék vége"},
		}),
	),
})
