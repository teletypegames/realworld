package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = ScriptManager.Register(Script{
	Name: inc.ScriptBVKLogged,
	Actions: inkwell.Seq(
		inkwell.Wait(0.4),
		inkwell.Say(inc.CandyCrusher, inkwell.Text{
			{Lang: "en", Text: "Why is my game beeping. It never beeps. Oh no. Oh no, that's a compliance screen."},
			{Lang: "hu", Text: "Miért sípol a játékom. Sose sípol. Ó, ne. Ó, ne, ez egy megfelelőségi képernyő."},
		}),
		inkwell.Say(inc.CandyCrusher, inkwell.Text{
			{Lang: "en", Text: "“Operator of record to present at BVK branch within the hour.” That's me. That's my licence. What did you type?"},
			{Lang: "hu", Text: "„A nyilvántartott operátor egy órán belül jelenjen meg a BVK-fiókban.” Ez én vagyok. Ez az én licenszem. Mit gépeltél be?"},
		}),
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Nothing you'd remember."},
			{Lang: "hu", Text: "Semmit, amire emlékeznél."},
		}),
		inkwell.Say(inc.CandyCrusher, inkwell.Text{
			{Lang: "en", Text: "Ninety-one levels. Ninety-one. Get away from me."},
			{Lang: "hu", Text: "Kilencvenegy szint. Kilencvenegy. Menj innen."},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "He's leaving in a hurry, and they'll be here in less of one. Move along, Paul. You've got what you came for."},
			{Lang: "hu", Text: "Ő sietve megy, ők még sietősebben jönnek. Lépj tovább, Paul. Megvan, amiért jöttél."},
		}),
		inkwell.SetFlag(inc.FlagBVKLogged),
	),
})
