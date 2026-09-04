package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var TerminalSearch = ScriptManager.Register(&Script{
	Name: "terminal_search",
	Actions: inkwell.Seq(
		character.Paul.Say(Text{
			{Lang: "en", Text: "“Workshop of the underwater sun.” Search."},
			{Lang: "hu", Text: "„A vízalatti nap műhelye.” Keresés."},
		}),
		inkwell.Wait(0.8),
		character.Paul.Say(Text{
			{Lang: "en", Text: "One hit. An anti-ACP fringe cult. They held a talk two weeks ago, in a club near an eatery in Chinatown."},
			{Lang: "hu", Text: "Egy találat. Egy ACP-ellenes bolondkultusz. Két héttel ezelőtt tartottak előadást, egy klubban, egy kínai negyedbeli kajálda mellett."},
		}),
		character.Paul.Say(Text{
			{Lang: "en", Text: "No exact address. “Location known to attendees.”"},
			{Lang: "hu", Text: "Pontos cím nincs. „A helyszínt a résztvevők ismerik.”"},
		}),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "Chinatown. A club nobody can find, next to a place everybody can. That's a start."},
			{Lang: "hu", Text: "Kínai negyed. Egy klub, amit senki nem talál, egy hely mellett, amit mindenki. Ez már valami."},
		}),
		inkwell.Wait(0.6),
		character.Paul.Say(Text{
			{Lang: "en", Text: "“This query has been logged by the BVK for operator review.” Wonderful."},
			{Lang: "hu", Text: "„Ezt a lekérdezést a BVK operátori felülvizsgálatra logolta.” Csodás."},
		}),
		inkwell.SetFlag(inc.FlagTerminalSearched),
		inkwell.RunScript(BVKLogged),
	),
})
