package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = ScriptManager.Register(Script{
	Name: inc.ScriptNormanTerminal,
	Actions: inkwell.Seq(
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Mail first. Half of it is about the simulation. “It could serve people in a way nothing has before.”"},
			{Lang: "hu", Text: "Először a levelek. A fele a szimulációról szól. „Úgy szolgálhatja az embereket, ahogy eddig semmi.”"},
		}),
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "The other half is about himself. “I don't understand why they still keep me on. What do I add to any of this?”"},
			{Lang: "hu", Text: "A másik fele magáról. „Nem értem, miért tartanak még. Mit teszek én hozzá ehhez az egészhez?”"},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "Lead engineer on the thing, and he writes like an intern who expects to be found out."},
			{Lang: "hu", Text: "Vezető mérnök a dolgon, és úgy ír, mint egy gyakornok, aki várja, hogy leleplezzék."},
		}),
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "The diary. Same note over and over: “My tape is smarter than I am. I'm only here because of it.”"},
			{Lang: "hu", Text: "A napló. Ugyanaz újra és újra: „A kazettám okosabb nálam. Csak miatta vagyok itt.”"},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "A man who thinks his tape carries him. I'd have words with that tape."},
			{Lang: "hu", Text: "Egy ember, aki szerint a kazettája viszi. Én beszélnék azzal a kazettával."},
		}),
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Noodle is all over this. The interface was his idea. He kept pushing Norman to take it further, into the trance states."},
			{Lang: "hu", Text: "Noodle mindenhol ott van. Az interfész az ő ötlete volt. Folyton hajtotta Normant, hogy vigye tovább, a transzállapotokig."},
		}),
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Last entry. He's complaining to his own tape that he and Noodle haven't had a beer in ages. That's all he wants. A beer, like two normal people."},
			{Lang: "hu", Text: "Utolsó bejegyzés. A saját kazettájának panaszkodik, hogy Noodle-lal ezer éve nem söröztek. Csak ezt akarja. Egy sört, mint két normális ember."},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "He wrote that to a tape. Then presumably the tape went to see Noodle. Think about the order of those two things."},
			{Lang: "hu", Text: "Ezt egy kazettának írta. Aztán a kazetta feltehetően elment Noodle-hoz. Gondold végig ennek a kettőnek a sorrendjét."},
		}),
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Practical things, too. A note to himself: the lift-room code at work. Written down. Of course it is."},
			{Lang: "hu", Text: "Gyakorlati dolgok is. Emlékeztető magának: a munkahelyi liftszoba kódja. Felírva. Hát persze."},
		}),
		inkwell.SetFlag(inc.FlagLiftCode),
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "And an open VPN session that timed out. It wants a login I don't have."},
			{Lang: "hu", Text: "És egy nyitott VPN-kapcsolat, ami lejárt. Belépőt kér, ami nincs nálam."},
		}),
		inkwell.SetFlag(inc.FlagVPNKnown),
		inkwell.If(inkwell.Flag(inc.FlagOfficeLogin),
			inkwell.Seq(
				inkwell.Say(inc.Paul, inkwell.Text{
					{Lang: "en", Text: "Wait. The sticky note from his desk at work. Trying it."},
					{Lang: "hu", Text: "Várj. A cetli a munkahelyi asztaláról. Kipróbálom."},
				}),
				inkwell.Say(inc.Dex, inkwell.Text{
					{Lang: "en", Text: "It's in. The tunnel ends at an office in the same complex. One that hasn't drawn power in a year."},
					{Lang: "hu", Text: "Bent van. Az alagút egy irodában ér véget, ugyanabban a komplexumban. Egy olyanban, ami egy éve nem vett fel áramot."},
				}),
				inkwell.SetFlag(inc.FlagAbandonedKnown),
			),
			inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "He'll have the login where he worked. People like him keep the password next to the machine that needs it."},
				{Lang: "hu", Text: "A belépő ott lesz, ahol dolgozott. Az ilyenek a gép mellett tartják a jelszót, amelyikhez kell."},
			}),
		),
		inkwell.SetFlag(inc.FlagNormanTerminalRead),
	),
})
