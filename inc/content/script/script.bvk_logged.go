package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var BVKLogged = ScriptManager.Register(&Script{
	Name: "bvk_logged",
	Actions: inkwell.Seq(
		inkwell.Wait(0.4),
		inkwell.Say(character.CandyCrusher, inkwell.Text{
			{Lang: "en", Text: "Why is my game beeping. It never beeps. Oh no. Oh no, that's a compliance screen."},
			{Lang: "hu", Text: "Miért sípol a játékom. Sose sípol. Ó, ne. Ó, ne, ez egy megfelelőségi képernyő."},
		}),
		inkwell.Say(character.CandyCrusher, inkwell.Text{
			{Lang: "en", Text: "“Operator of record to present at BVK branch within the hour.” That's me. That's my licence. What did you type?"},
			{Lang: "hu", Text: "„A nyilvántartott operátor egy órán belül jelenjen meg a BVK-fiókban.” Ez én vagyok. Ez az én licenszem. Mit gépeltél be?"},
		}),
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "Nothing you'd remember."},
			{Lang: "hu", Text: "Semmit, amire emlékeznél."},
		}),
		inkwell.Say(character.CandyCrusher, inkwell.Text{
			{Lang: "en", Text: "Ninety-one levels. Ninety-one. Get away from me."},
			{Lang: "hu", Text: "Kilencvenegy szint. Kilencvenegy. Menj innen."},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "He's leaving in a hurry, and they'll be here in less of one. Move along, Paul. You've got what you came for."},
			{Lang: "hu", Text: "Ő sietve megy, ők még sietősebben jönnek. Lépj tovább, Paul. Megvan, amiért jöttél."},
		}),
		inkwell.SetFlag(inc.FlagBVKLogged),
	),
})
