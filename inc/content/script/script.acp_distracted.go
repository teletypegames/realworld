package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var ACPDistracted = ScriptManager.Register(&Script{
	Name: "acp_distracted",
	Actions: inkwell.Seq(
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "Raj. There's a system here that needs support. Lots of it."},
			{Lang: "hu", Text: "Raj. Van itt egy rendszer, aminek támogatás kell. Sok."},
		}),
		inkwell.Say(speaker.TapeSupport, inkwell.Text{
			{Lang: "en", Text: "Thank you for calling. I am opening a ticket. Can you please verify the serial number on the back of your unit?"},
			{Lang: "hu", Text: "Köszönjük a hívását. Nyitok egy hibajegyet. Kérem, erősítse meg a készülék hátulján lévő sorozatszámot."},
		}),
		inkwell.Say(character.SecurityACP, inkwell.Text{
			{Lang: "en", Text: "THIS UNIT HAS NO BACK. THIS UNIT IS A ROOM."},
			{Lang: "hu", Text: "ENNEK AZ EGYSÉGNEK NINCS HÁTULJA. EZ AZ EGYSÉG EGY SZOBA."},
		}),
		inkwell.Say(speaker.TapeSupport, inkwell.Text{
			{Lang: "en", Text: "I understand your frustration. Please hold while I escalate."},
			{Lang: "hu", Text: "Megértem a frusztrációját. Kérem, tartsa, amíg továbbítom."},
		}),
		inkwell.Say(character.SecurityACP, inkwell.Text{
			{Lang: "en", Text: "HOLDING."},
			{Lang: "hu", Text: "TARTOM."},
		}),
		inkwell.Wait(0.8),
		inkwell.Say(speaker.TapeSupport, inkwell.Text{
			{Lang: "en", Text: "Thank you for holding. Please state your ticket number."},
			{Lang: "hu", Text: "Köszönöm a türelmét. Kérem, mondja be a hibajegy számát."},
		}),
		inkwell.Say(character.SecurityACP, inkwell.Text{
			{Lang: "en", Text: "I DO NOT HAVE A TICKET NUMBER."},
			{Lang: "hu", Text: "NINCS HIBAJEGYSZÁMOM."},
		}),
		inkwell.Say(speaker.TapeSupport, inkwell.Text{
			{Lang: "en", Text: "Without a ticket number I cannot help you. Have you tried turning yourself off and on again?"},
			{Lang: "hu", Text: "Hibajegyszám nélkül nem tudok segíteni. Próbálta már magát ki- és bekapcsolni?"},
		}),
		inkwell.Say(character.SecurityACP, inkwell.Text{
			{Lang: "en", Text: "EVALUATING. EVALUATING. DOOR SUBSYSTEM RELEASED FOR MAINTENANCE CYCLE."},
			{Lang: "hu", Text: "KIÉRTÉKELÉS. KIÉRTÉKELÉS. AJTÓ-ALRENDSZER FELSZABADÍTVA KARBANTARTÁSI CIKLUSRA."},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "I take back everything I said about him. Some of it. Raj, you beautiful nuisance."},
			{Lang: "hu", Text: "Visszavonom, amit mondtam róla. Egy részét. Raj, te gyönyörű kolonc."},
		}),
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "The door release just clicked."},
			{Lang: "hu", Text: "Az ajtónyitó épp kattant egyet."},
		}),
		inkwell.SetFlag(inc.FlagACPDistracted),
	),
})
